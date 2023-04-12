package com.ams.backend.configuation

import com.ams.backend.filter.JwtAuthorizationFilter
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.HttpStatus
import org.springframework.security.authentication.AuthenticationManager
import org.springframework.security.config.annotation.web.builders.HttpSecurity
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity
import org.springframework.security.config.annotation.web.invoke
import org.springframework.security.web.authentication.HttpStatusEntryPoint
import org.springframework.security.config.annotation.authentication.configuration.AuthenticationConfiguration
import org.springframework.security.web.SecurityFilterChain

@Configuration
@EnableWebSecurity
class SecurityConfiguration(val authenticationConfiguration: AuthenticationConfiguration) {
    private val IGNORE_API = arrayOf("/admin/login", "/admin/list", "/admin/register")


    fun authenticationManager(): AuthenticationManager =
        authenticationConfiguration.authenticationManager


    @Bean
    fun filterChain(http: HttpSecurity): SecurityFilterChain {
        http {
            cors { }
            csrf { disable() }

            authorizeHttpRequests {
                IGNORE_API.forEach { ignoreApi ->
                    authorize(ignoreApi, permitAll)
                }
                authorize("/**", authenticated)
            }
            httpBasic {}
//            authenticationManagerBuilder.userDetailsService(adminService)

            exceptionHandling {
                // Send 401 response when the user does not have access to the resource
//                authenticationEntryPoint = HttpStatusEntryPoint(HttpStatus.UNAUTHORIZED)
            }
            apply(securityConfigurationAdapter())
        }
        return http.build()
    }


    private fun securityConfigurationAdapter() =
        JwtConfigurer(JwtAuthorizationFilter(authenticationManager()))

}