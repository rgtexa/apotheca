# apotheca
A document repository written in Go

# Role-based Permissions

- Read: 1
- Update Owned: 2
- Create: 4
- Delete Owned: 8
- Manage (CRUD) Department Docs and Routes: 16
- Manage (CRUD) Global Docs and Routes: 32
- Full Perms: 64

## Roles Explanation

### System Admin

System Admins will have full permissions to CRUD documents and routes, and manage user permissions (perm = 127).

### Document Admins

Document Admins will have permissions to CRUD documents and routes, but not user permissions (perm = 63).

### Department Managers

Managers will have permissions to CRUD documents and routes belonging to their department (perm = 31).

### Authors

Authors will have permissions to CRUD documents where they are owners (perm = 15).

### Users

Users have basic Read-only access to documents (perm = 1).
