package com.valet.gurren.repo.db.userRepo

import com.valet.gurren.model.User
import org.springframework.data.jpa.repository.Modifying
import org.springframework.data.jpa.repository.Query
import org.springframework.data.repository.CrudRepository
import org.springframework.data.repository.query.Param

interface UserRepo : CrudRepository<User, String> {
    @Modifying
    @Query("UPDATE User u set u.isPremium = :status where u.username = :username")
    fun updateUserPremiumStatus(@Param("username") username: String, @Param("status") status: Boolean)

    @Modifying
    @Query("update User u set u.isEnable = false where u.username = :username")
    fun deleteUser(@Param("username") username: String)

    @Modifying
    @Query("update User u set u.isEnable = true where u.username = :username")
    fun recoveryUser(@Param("username") username: String)
}