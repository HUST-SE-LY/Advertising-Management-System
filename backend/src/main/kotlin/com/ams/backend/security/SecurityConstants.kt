package com.ams.backend.security

object SecurityConstants {

    // from https://www.allkeysgenerator.com/
    // TODO(Write the key into the .property configuration file)
    const val JWT_SECRET_KEY = "5ee13f86d7b149f386ef49fbb37fb62b"

    const val TOKEN_TYPE = "JWT"
    const val TOKEN_HEADER = "Authorization"

    /**
     * Generally, Authorization is added to the request header and Bearer tag is added.
     */
    const val TOKEN_PREFIX = "Bearer "

    const val TOKEN_ROLE_CLAIM = "role"
    const val TOKEN_ISSUER = "security"
    const val TOKEN_AUDIENCE = "security-all"

    /**
     * When Remember is false, the token is valid for 2 hours
     */
    const val EXPIRATION_TIME = 60 * 60 * 2L

    /**
     * When Remember is true, the token is valid for 7 days
     */
    const val EXPIRATION_REMEMBER_TIME = 60 * 60 * 24 * 7L
}