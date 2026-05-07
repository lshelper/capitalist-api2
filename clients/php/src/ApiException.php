<?php

namespace Capitalist\Api2;

use RuntimeException;

class ApiException extends RuntimeException
{
    private $statusCode;
    private $responseBody;

    public function __construct($message, $statusCode = 0, $responseBody = '')
    {
        parent::__construct($message, $statusCode);
        $this->statusCode = $statusCode;
        $this->responseBody = $responseBody;
    }

    public function getStatusCode()
    {
        return $this->statusCode;
    }

    public function getResponseBody()
    {
        return $this->responseBody;
    }
}
