package com.ams.backend.filter

import com.ams.backend.security.SecurityConstants
import com.ams.backend.utils.JwtTokenUtil
import jakarta.servlet.FilterChain
import jakarta.servlet.http.HttpServletRequest
import jakarta.servlet.http.HttpServletResponse
import org.springframework.security.authentication.AuthenticationManager
import org.springframework.security.core.context.SecurityContextHolder
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter
import org.springframework.util.StringUtils

class JwtAuthorizationFilter(authenticationManager: AuthenticationManager): BasicAuthenticationFilter(authenticationManager) {

    override fun doFilterInternal(request: HttpServletRequest, response: HttpServletResponse, chain: FilterChain) {
        val token = getTokenFromHttpRequest(request)
        if (!token.isNullOrBlank() && JwtTokenUtil.validateToken(token)) {
            val authentication = JwtTokenUtil.getAuthentication(token)
            // Save authentication information into the Spring security context
            SecurityContextHolder.getContext().authentication = authentication
        }
        chain.doFilter(request, response)
    }


    private fun getTokenFromHttpRequest(request: HttpServletRequest): String? {
        val authorization = request.getHeader(SecurityConstants.TOKEN_HEADER) ?: return null
        if (!authorization.startsWith(SecurityConstants.TOKEN_ISSUER)) {
            return null
        }

        // Remove prefix
        return authorization.replace(SecurityConstants.TOKEN_PREFIX, "")

    }
}