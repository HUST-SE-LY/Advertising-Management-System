package com.ams.backend.admin.model

import com.ams.backend.GlobalConstants
import com.ams.backend.admin.dto.AdminDTO
import org.springframework.data.annotation.Id
import org.springframework.data.relational.core.mapping.Table
import org.springframework.security.core.userdetails.User

@Table(name = "Admin", schema = GlobalConstants.SCHEME)
data class Admin(@Id val id: Long? = null, val name: String, val password: String) {

    fun toUser(): User = User(name, password, listOf(AdminAuthority))

    fun toAdminDTO() = AdminDTO(name = name)

    override fun equals(other: Any?) = other is Admin && id == other.id

    override fun hashCode() = id.hashCode()

//    override fun getAuthority(): String = SecurityConstants.ROLE_ADMIN
}


