package com.ams.security

import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.security.config.annotation.web.reactive.EnableWebFluxSecurity
import org.springframework.security.config.web.server.ServerHttpSecurity
import org.springframework.security.config.web.server.invoke
import org.springframework.security.web.server.SecurityWebFilterChain
import org.springframework.security.web.server.context.NoOpServerSecurityContextRepository
import org.springframework.web.reactive.config.EnableWebFlux


@Configuration
@EnableWebFlux
@EnableWebFluxSecurity
class SecurityConfiguration {
    private val IGNORE_API = arrayOf("/admin/login")


    @Bean
    fun filterChain(
        http: ServerHttpSecurity,
    ): SecurityWebFilterChain =
        http {
            cors {
            }
            csrf { disable() }
            authorizeExchange {
                authorize("/**", permitAll)
//                authorize("/admin/register", authenticated)
//                authorize("/customer/**", hasRole(SecurityConstants.ROLE_CUSTOMER))
                IGNORE_API.forEach { ignoreApi ->
                    authorize(ignoreApi, permitAll)
                }
//                authorize(anyExchange, authenticated)
            }

            httpBasic {
                securityContextRepository = NoOpServerSecurityContextRepository.getInstance()
            }

//            headers {
//                cache { disable() }
//            }
        }
}