package com.ams.backend.security


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
//        @Qualifier("ADMIN") jwtAdminAuthenticationFilter: AuthenticationWebFilter,
        jwtService: JwtService
    ): SecurityWebFilterChain =
        http {
            cors {
//                configurationSource = TODO(Add front-end site)
            }
            csrf { disable() }
            authorizeExchange {
                authorize("/**", permitAll)
//                authorize("/admin/register", authenticated)
//                authorize("/customer/**", hasRole(SecurityConstants.ROLE_CUSTOMER))
//                IGNORE_API.forEach { ignoreApi ->
//                    authorize(ignoreApi, permitAll)
//                }
//                authorize(anyExchange, authenticated)
            }

            httpBasic {
                securityContextRepository = NoOpServerSecurityContextRepository.getInstance()
            }

//            headers {
//                cache { disable() }
//            }
        }

//    @Bean
//    @Qualifier("ADMIN")
//    fun authenticationWebFilter(
//        reactiveAuthenticationManager: ReactiveAuthenticationManager,
//        jwtConverter: ServerAuthenticationConverter,
//        serverAuthenticationSuccessHandler: ServerAuthenticationSuccessHandler,
//        jwtServerAuthenticationFailureHandler: ServerAuthenticationFailureHandler
//    ): AuthenticationWebFilter {
//        return AuthenticationWebFilter(reactiveAuthenticationManager).apply {
//            setRequiresAuthenticationMatcher {
//                ServerWebExchangeMatchers.pathMatchers(
//                    HttpMethod.POST,
//                    "/admin/login"
//                ).matches(it)
//            }
//            setServerAuthenticationConverter(jwtConverter)
//            setAuthenticationSuccessHandler(serverAuthenticationSuccessHandler)
//            setAuthenticationFailureHandler(jwtServerAuthenticationFailureHandler)
//            setSecurityContextRepository(NoOpServerSecurityContextRepository.getInstance())
//        }
//    }
//
//    @Bean
//    fun reactiveAuthenticationManager(
//        reactiveUserDetailsService: AdminService,
//        passwordEncoder: PasswordEncoder
//    ): ReactiveAuthenticationManager {
//        val manager = UserDetailsRepositoryReactiveAuthenticationManager(reactiveUserDetailsService)
//        manager.setPasswordEncoder(passwordEncoder)
//        return manager
//    }
}