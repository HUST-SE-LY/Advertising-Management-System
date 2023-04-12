package com.ams.backend.configuation

import com.ams.backend.filter.JwtAuthorizationFilter
import org.springframework.security.config.annotation.SecurityConfigurerAdapter
import org.springframework.security.config.annotation.web.builders.HttpSecurity
import org.springframework.security.web.DefaultSecurityFilterChain
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter

class JwtConfigurer(private val jwtAuthorizationFilter: JwtAuthorizationFilter):
    SecurityConfigurerAdapter<DefaultSecurityFilterChain, HttpSecurity>() {

    override fun configure(http: HttpSecurity) {
        http.addFilterBefore(jwtAuthorizationFilter, UsernamePasswordAuthenticationFilter::class.java)
    }


}