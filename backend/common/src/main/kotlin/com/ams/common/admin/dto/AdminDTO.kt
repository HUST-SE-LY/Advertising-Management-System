package com.ams.backend.admin.dto

import kotlinx.serialization.Serializable

@Serializable
data class AdminDTO(val name: String = "", val errorMessage: String? = null)