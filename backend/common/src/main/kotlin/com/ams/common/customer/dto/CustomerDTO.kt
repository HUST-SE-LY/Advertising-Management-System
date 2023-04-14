package com.ams.backend.customer.dto

import kotlinx.serialization.Serializable

@Serializable
data class CustomerDTO(
    val account: String,
    val name: String? = null,
    val address: String? = null,
    val managerName: String? = null,
    val managerTel: String? = null,
    val businessLicenseNumber: String? = null
)