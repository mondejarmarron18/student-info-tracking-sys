
-- ROLE TABLE
INSERT INTO role (name, description)
VALUES 
('admin', 'Administrator role with full system access, including user management, system settings, and advanced features.'),
('student', 'Student role with restricted access to personal information update, and messaging features.'),
('teacher', 'Staff role with permissions to manage student update request, evaluate the student info authenticity, and update student login access.');

-- PERMISSION TABLE
INSERT INTO permission (name, description)
VALUES
-- User Account Permissions
('view:userAccounts', 'Permission to view information for all user accounts in the system.'),
('view:userAccount', 'Permission to view detailed information about a specific user account.'),
('view:ownUserAccount', 'Permission to view information related to the currently logged-in user account.'),
('create:userAccount', 'Permission to create new user accounts with associated profiles and roles.'),
('create:ownUserAccount', 'Permission to create and manage the logged-in user''s own account details.'),
('update:ownUserAccount', 'Permission to update the logged-in user''s own account details and settings.'),
('update:userAccount', 'Permission to update details of specific user accounts, including roles or status.'),
('delete:userAccount', 'Permission to delete specific user accounts from the system.'),

-- Course Management Permissions
('view:courses', 'Permission to view courses available in the system.'),
('view:course', 'Permission to view detailed information about a specific course.'),
('create:course', 'Permission to create new courses in the system.'),
('update:course', 'Permission to update details of existing courses, such as descriptions or levels.'),
('delete:course', 'Permission to remove a course from the system.'),

-- Course Level Permissions
('view:courseLevels', 'Permission to view course levels associated with a specific course or student.'),
('create:courseLevel', 'Permission to create new course levels for a course.'),
('update:courseLevel', 'Permission to update existing course level details.'),
('delete:courseLevel', 'Permission to delete a course level associated with a course.'),

-- Course Specialization Permissions
('view:courseSpecializations', 'Permission to view specializations associated with a course level.'),
('create:courseSpecialization', 'Permission to create a new specialization for a course level.'),
('update:courseSpecialization', 'Permission to update the details of an existing course specialization.'),
('delete:courseSpecialization', 'Permission to delete a specialization associated with a course level.'),

-- Student and Guardian Permissions
('view:studentGuardians', 'Permission to view information about student guardians.'),
('create:studentGuardian', 'Permission to create a new guardian record for a student.'),
('update:studentGuardian', 'Permission to update details of an existing guardian record.'),
('delete:studentGuardian', 'Permission to delete a guardian record associated with a student.'),

-- Approval and Audit Permissions
('view:updateApprovals', 'Permission to view update approvals in the system.'),
('update:updateApproval', 'Permission to review and approve or reject pending update request.'),
('delete:updateApproval', 'Permission to delete an update approval record.');
