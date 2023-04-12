package com.ams.backend.service

import com.ams.backend.model.entity.Admin
import com.ams.backend.repository.AdminRepository
import kotlinx.coroutines.flow.toList
import org.springframework.security.core.userdetails.UsernameNotFoundException
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.stereotype.Service

@Service
class AdminService(private val adminRepository: AdminRepository, val bCryptPasswordEncoder: BCryptPasswordEncoder) {


    suspend fun registerAdmin(name: String, password: String) =
        run {
            val admin = Admin(name = name, password = bCryptPasswordEncoder.encode(password))
            adminRepository.save(admin)
        }

//    suspend fun login(admin: Admin): Admin {
//        val upaToken = UsernamePasswordAuthenticationToken(admin.name, admin.password)
//
//        val id = admin.id ?: throw IllegalArgumentException("The id of admin shouldn't be null")
//        val adminInDatabase =
//            adminRepository.findById(id) ?: throw EmptyResultDataAccessException("Excepted one result, got nothing", 1)
//        if (admin.name != adminInDatabase.name) throw IllegalArgumentException("The name of admin isn't matched")
//        if (!PasswordUtil.isPasswordMatch(rawPassword = admin.password, encodedPassword = adminInDatabase.password)) {
//            throw IllegalArgumentException("The password isn't correct")
//        }
//        return adminInDatabase
//    }

    suspend fun getAdminByName(name: String): Admin =
        adminRepository.findAdminByName(name) ?: throw UsernameNotFoundException("Admin not found with username: $name")


    suspend fun listAllAdmins(): List<Admin> =
        adminRepository.findAll().toList()


}