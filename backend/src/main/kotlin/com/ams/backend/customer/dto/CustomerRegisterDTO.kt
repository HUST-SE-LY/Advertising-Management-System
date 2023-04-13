package com.ams.backend.customer.dto

data class CustomerRegisterDTO(
    val name: String,
    val account: String,
    val password: String,
    val address: String,
    val managerName: String,
    val managerTel: String,
    val businessLicenseNumber: String
)