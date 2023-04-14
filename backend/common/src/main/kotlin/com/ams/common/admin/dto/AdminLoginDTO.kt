package com.ams.backend.admin.dto

import kotlinx.serialization.Serializable

@Serializable
data class AdminLoginDTO(val name: String, val password: String)
