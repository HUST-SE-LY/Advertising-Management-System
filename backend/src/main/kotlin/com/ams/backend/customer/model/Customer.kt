package com.ams.backend.customer.model

import org.springframework.data.relational.core.mapping.Table
import org.springframework.security.core.userdetails.User

@Table("customer", schema = "ams")
data class Customer(val name: String, val account: String, val password: String,val address: String, val managerName: String, val managerTel: String, val businessLicenseNumber: String) {

    fun toUser() = User(name, password, listOf(CustomerAuthority))

}
