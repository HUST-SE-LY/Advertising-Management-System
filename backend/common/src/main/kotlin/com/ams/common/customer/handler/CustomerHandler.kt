package com.ams.backend.customer.handler

import com.ams.backend.customer.dto.CustomerRegisterDTO
import com.ams.backend.customer.model.Customer
import com.ams.backend.customer.repository.CustomerRepository
import com.ams.backend.security.JwtService
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import org.springframework.http.HttpHeaders
import org.springframework.http.HttpStatus
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.*
import kotlin.time.Duration.Companion.days
import kotlin.time.Duration.Companion.hours

@Component
class CustomerHandler(
    private val customerRepository: CustomerRepository,
    private val jwtService: JwtService,
    private val bCryptPasswordEncoder: BCryptPasswordEncoder
) {

    companion object {
        private val ACCESS_TOKEN_EXPIRATION = 2.hours
        private val REFRESH_TOKEN_EXPIRATION = 7.days
    }

    suspend fun register(request: ServerRequest): ServerResponse {
        val body = request.awaitBody<String>()
        val customerRegisterDTO = Json.decodeFromString<CustomerRegisterDTO>(body)

        if (customerRepository.findCustomerByAccount(customerRegisterDTO.account) != null) {
            return ServerResponse.status(HttpStatus.CONFLICT).bodyValueAndAwait("The same account name already exists")
        }

        val encryptedPassword = bCryptPasswordEncoder.encode(customerRegisterDTO.password)
        val customer = Customer(
            id = null,
            account = customerRegisterDTO.account,
            password = encryptedPassword,
            name = customerRegisterDTO.name,
            address = customerRegisterDTO.address,
            managerName = customerRegisterDTO.managerName,
            managerTel = customerRegisterDTO.managerTel,
            businessLicenseNumber = customerRegisterDTO.businessLicenseNumber
        )

        customerRepository.save(customer)

        val (accessToken, refreshToken) = jwtService.generateTokenFromUser(customer.toUser())

        return ServerResponse.ok().headers {
            it[HttpHeaders.AUTHORIZATION] = accessToken
        }.bodyValueAndAwait(Json.encodeToString(customer.toCustomerDTO()))
    }


}