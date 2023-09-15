SELECT account.*
FROM account
         LEFT JOIN account_role ON account.id = account_role.AccountId
         LEFT JOIN role ON role.id = account_role.RoleId
WHERE RoleId = ?
  AND account.deleted = 0
  AND account_role.deleted = 0
  AND role.deleted = 0;

SELECT role.*
FROM role
         LEFT JOIN account_role ON role.id = account_role.RoleId
         LEFT JOIN account ON account.id = account_role.AccountId
WHERE AccountId = ?
  AND role.deleted = 0
  AND account_role.deleted = 0
  AND account.deleted = 0;