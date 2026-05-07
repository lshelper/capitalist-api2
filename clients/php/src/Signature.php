<?php

namespace Capitalist\Api2;

final class Signature
{
    public static function calculate($timestamp, $body, $apiSecret)
    {
        return hash('sha256', $timestamp . $body . $apiSecret);
    }

    public static function verify($timestamp, $body, $apiSecret, $signature)
    {
        return hash_equals(self::calculate($timestamp, $body, $apiSecret), $signature);
    }
}
