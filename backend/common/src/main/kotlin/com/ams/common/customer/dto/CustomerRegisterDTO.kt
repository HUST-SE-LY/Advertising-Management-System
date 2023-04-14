package com.ams.backend.customer.dto

import kotlinx.serialization.Serializable

@Serializable
data class CustomerRegisterDTO(
    val account: String,
    val password: String,
    val name: String? = null,
    val address: String? = null,
    val managerName: String? = null,
    val managerTel: String? = null,
    val businessLicenseNumber: String? = null
)