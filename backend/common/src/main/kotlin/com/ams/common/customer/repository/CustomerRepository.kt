package com.ams.backend.customer.repository

import com.ams.backend.customer.model.Customer
import org.springframework.data.repository.kotlin.CoroutineCrudRepository

interface CustomerRepository: CoroutineCrudRepository<Customer, Long> {
    suspend fun findCustomerByAccount(account: String): Customer?
}