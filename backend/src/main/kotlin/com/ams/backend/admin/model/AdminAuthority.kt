package com.ams.backend.admin.model

import com.ams.backend.security.SecurityConstants
import org.springframework.security.core.GrantedAuthority

object AdminAuthority: GrantedAuthority {
    override fun getAuthority(): String = SecurityConstants.ROLE_ADMIN
}