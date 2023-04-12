package com.ams.backend.model.entity

import kotlinx.serialization.Serializable
import org.springframework.data.annotation.Id
import org.springframework.data.relational.core.mapping.Column
import org.springframework.data.relational.core.mapping.Table

@Serializable
@Table(name = "Admin", schema = "ams")
data class Admin(@Id val id: Long? = null, @Column("name") val name: String, val password: String)
