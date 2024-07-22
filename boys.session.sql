
           SELECT 
            User.id AS id,
        User.name AS name,
        User.email AS email,
        User.mobile AS mobile,
        User.roleId AS roleId,
        User.webToken AS webToken,
        User.isActive AS isActive,
        Role.name AS roleName,
        UserLoginRequest.createdAt AS otpCreatedAt
            FROM User
       INNER JOIN Role ON Role.id = User.roleId
       INNER JOIN UserLoginRequest ON UserLoginRequest.userId = User.id
     --           INNER JOIN LatestRequests ON UserLoginRequest.userId = LatestRequests.userId AND UserLoginRequest.createdAt = LatestRequests.latestCreatedAt
            WHERE User.email = "venukumar.rvk@gmail.com" AND User.isActive = 1;