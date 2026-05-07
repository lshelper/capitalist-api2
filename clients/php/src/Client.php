<?php

namespace Capitalist\Api2;

final class Client
{
    const DEFAULT_BASE_URL = 'https://api2.capitalist.net';

    private $apiKey;
    private $apiSecret;
    private $baseUrl;

    public function __construct($apiKey, $apiSecret, $baseUrl = self::DEFAULT_BASE_URL)
    {
        $this->apiKey = $apiKey;
        $this->apiSecret = $apiSecret;
        $this->baseUrl = rtrim($baseUrl, '/');
    }

    public function readWhitelist()
    {
        return $this->request('GET', '/v1/whitelist');
    }

    public function addWhitelist(array $ip)
    {
        return $this->request('POST', '/v1/whitelist', array('ip' => array_values($ip)));
    }

    public function removeWhitelist(array $ip)
    {
        return $this->request('POST', '/v1/whitelist/remove', array('ip' => array_values($ip)));
    }

    public function listAccounts($currency)
    {
        return $this->request('GET', '/v1/account/list', null, array('currency' => $currency));
    }

    public function createExchange($fromAccount, $toAccount, $amount)
    {
        return $this->request('POST', '/v1/exchange', array(
            'fromAccount' => $fromAccount,
            'toAccount' => $toAccount,
            'amount' => $amount,
        ));
    }

    public function getRate($from, $to)
    {
        return $this->request('GET', '/v1/rate', null, array('from' => $from, 'to' => $to));
    }

    public function createPayment(array $payment)
    {
        return $this->request('POST', '/v1/payment', $payment);
    }

    public function getPaymentByDocumentId($documentId)
    {
        return $this->request('GET', '/v1/payment/document/' . rawurlencode((string) $documentId));
    }

    public function getPaymentByUserRequestId($userRequestId)
    {
        return $this->request('GET', '/v1/payment/' . rawurlencode((string) $userRequestId));
    }

    public function getOrders(array $filters = array())
    {
        return $this->request('GET', '/v1/orders', null, $filters);
    }

    public function getTransactions(array $filters = array())
    {
        return $this->request('GET', '/v1/transactions', null, $filters);
    }

    public function startKyc(array $request)
    {
        return $this->request('POST', '/v1/kyc/start', $request);
    }

    public function getKycStatus($kycExternalUserId, $sort)
    {
        return $this->request('GET', '/v1/kyc/status/' . rawurlencode($kycExternalUserId) . '/' . rawurlencode($sort));
    }

    public function getKycStatusByUuid($uuid)
    {
        return $this->request('GET', '/v1/kyc/statusByUuid/' . rawurlencode($uuid));
    }

    public function setKycData($uuid, array $data)
    {
        return $this->request('POST', '/v1/kyc/setData/' . rawurlencode($uuid), $data);
    }

    public function setKycPicture($uuid, $type, $jpegBytes)
    {
        return $this->requestRaw(
            'POST',
            '/v1/kyc/setPicture/' . rawurlencode($uuid) . '/' . rawurlencode($type),
            $jpegBytes,
            'image/jpeg'
        );
    }

    public function confirmKyc($uuid)
    {
        return $this->request('POST', '/v1/kyc/confirm/' . rawurlencode($uuid), new \stdClass());
    }

    private function request($method, $path, $body = null, array $query = array())
    {
        $rawBody = $body === null ? '' : json_encode($body, JSON_UNESCAPED_SLASHES);

        if ($rawBody === false) {
            throw new ApiException('Failed to encode request body as JSON');
        }

        return $this->requestRaw($method, $path, $rawBody, 'application/json', $query);
    }

    private function requestRaw($method, $path, $rawBody = '', $contentType = 'application/json', array $query = array())
    {
        $url = $this->baseUrl . $path;

        if ($query) {
            $url .= '?' . http_build_query($query, '', '&', PHP_QUERY_RFC3986);
        }

        $timestamp = (string) round(microtime(true) * 1000);
        $signature = Signature::calculate($timestamp, $rawBody, $this->apiSecret);
        $headers = array(
            'API-Key: ' . $this->apiKey,
            'X-Request-Timestamp: ' . $timestamp,
            'Signature: ' . $signature,
        );

        if ($rawBody !== '') {
            $headers[] = 'Content-Type: ' . $contentType;
        }

        $ch = curl_init($url);
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, $method);
        curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

        if ($rawBody !== '') {
            curl_setopt($ch, CURLOPT_POSTFIELDS, $rawBody);
        }

        $responseBody = curl_exec($ch);

        if ($responseBody === false) {
            $message = curl_error($ch);
            curl_close($ch);
            throw new ApiException($message);
        }

        $statusCode = (int) curl_getinfo($ch, CURLINFO_HTTP_CODE);
        curl_close($ch);

        if ($statusCode < 200 || $statusCode >= 300) {
            $message = $this->extractErrorMessage($responseBody);
            throw new ApiException($message, $statusCode, $responseBody);
        }

        if ($responseBody === '' || $responseBody === 'null') {
            return null;
        }

        $decoded = json_decode($responseBody, true);
        return json_last_error() === JSON_ERROR_NONE ? $decoded : $responseBody;
    }

    private function extractErrorMessage($responseBody)
    {
        $decoded = json_decode($responseBody, true);

        if (json_last_error() === JSON_ERROR_NONE && isset($decoded['error'])) {
            return (string) $decoded['error'];
        }

        return $responseBody ?: 'Capitalist API request failed';
    }
}
