package com.valet.gurren.repo.db.userSettingsProfileRepo

import com.valet.gurren.model.UserProfileSetting
import org.springframework.data.repository.CrudRepository

interface UserSettingsProfileRepo : CrudRepository<UserProfileSetting, String> {
}