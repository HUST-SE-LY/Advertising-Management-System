package com.ams.backend.dto

import kotlinx.serialization.Serializable

@Serializable
data class AdminDTO(val admin: String = "", val errorMessage: String? = null)