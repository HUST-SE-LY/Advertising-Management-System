package com.ams.backend.admin.repository

import com.ams.backend.admin.model.Admin
import org.springframework.data.repository.kotlin.CoroutineCrudRepository

interface AdminRepository : CoroutineCrudRepository<Admin, Long> {
    suspend fun findAdminByName(name: String): Admin?
}