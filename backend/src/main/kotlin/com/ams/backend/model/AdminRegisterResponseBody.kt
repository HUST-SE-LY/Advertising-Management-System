package com.ams.backend.model

import com.ams.backend.model.entity.Admin
import kotlinx.serialization.Serializable

@Serializable
data class AdminRegisterResponseBody(val admin: List<Admin> = emptyList(), val errorMessage: String? = null) {

    companion object {
        fun from(adminResult: Result<List<Admin>>): AdminRegisterResponseBody {
            val admin = adminResult.getOrElse { exception: Throwable ->
                return AdminRegisterResponseBody(errorMessage = exception.localizedMessage)
            }
            return AdminRegisterResponseBody(admin = admin, errorMessage = null)
        }
    }
}

//sealed interface AdminRegisterResponse {
//
//    fun toResponseEntity() = run {
//        val (admin, status) = if (this is Success) {
//            this.admin to HttpStatus.OK
//        } else {
//            null to HttpStatus.NOT_ACCEPTABLE
//        }
//        ResponseEntity<Admin>(admin, status)
//    }
//
//    class Success(val admin: Admin) : AdminRegisterResponse
//
//    class Failure(val message: String) : AdminRegisterResponse
//}