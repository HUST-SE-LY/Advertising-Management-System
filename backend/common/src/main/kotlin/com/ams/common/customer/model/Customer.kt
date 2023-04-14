package com.ams.backend.customer.model

import com.ams.backend.GlobalConstants
import com.ams.backend.customer.dto.CustomerDTO
import org.springframework.data.annotation.Id
import org.springframework.data.relational.core.mapping.Column
import org.springframework.data.relational.core.mapping.Table
import org.springframework.security.core.userdetails.User

@Table(name = "customer", schema = GlobalConstants.SCHEME)
data class Customer(
    @Id val id: Long? = null,
    val account: String,
    val password: String,
    val name: String? = null,
    val address: String? = null,
    @Column("manager_name") val managerName: String? = null,
    @Column("manager_tel") val managerTel: String? = null,
    @Column("business_license_number") val businessLicenseNumber: String? = null
) {
    fun toUser() = User(name, password, listOf(CustomerAuthority))

    fun toCustomerDTO() = CustomerDTO(account, name, address, managerName, managerTel, businessLicenseNumber)
}
