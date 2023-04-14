package com.ams.backend.customer.model

import com.ams.backend.security.SecurityConstants
import org.springframework.security.core.GrantedAuthority

object CustomerAuthority: GrantedAuthority {

    override fun getAuthority(): String = SecurityConstants.ROLE_CUSTOMER
}