package com.ams.backend.repository

import com.ams.backend.model.entity.Admin
import kotlinx.coroutines.flow.Flow
import org.springframework.data.r2dbc.repository.Query
import org.springframework.data.repository.kotlin.CoroutineCrudRepository

interface AdminRepository : CoroutineCrudRepository<Admin, Long> {
    suspend fun findAdminByName(name: String): Admin?
}