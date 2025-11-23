-- 创建数据库
CREATE DATABASE IF NOT EXISTS ark_iam DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
USE ark_iam;

-- 1. 租户核心表
-- ============================================

-- 租户表(最高层级)
CREATE TABLE iam_tenant (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '租户ID',
    tenant_code VARCHAR(32) UNIQUE NOT NULL COMMENT '租户编码',
    tenant_name VARCHAR(64) NOT NULL COMMENT '租户名称',
    description VARCHAR(500) COMMENT '租户描述',
    status VARCHAR(16) DEFAULT 'active' COMMENT '状态: active-正常 inactive-停用',
    sort_order INT DEFAULT 0 COMMENT '排序',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    INDEX idx_created_at (created_at),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='租户表';

-- 公司表(租户主体)
CREATE TABLE iam_company (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '公司ID(租户ID)',
    tenant_id BIGINT NOT NULL COMMENT '所属租户ID',
    company_code VARCHAR(32) NOT NULL COMMENT '公司编码',
    company_name VARCHAR(128) NOT NULL COMMENT '公司名称',
    short_name VARCHAR(64) COMMENT '公司简称',
    unified_social_credit_code VARCHAR(18) COMMENT '统一社会信用代码(18位)',
    legal_person VARCHAR(32) COMMENT '法人代表',
    contact_phone VARCHAR(16) COMMENT '联系电话',
    contact_email VARCHAR(64) COMMENT '联系邮箱',
    address VARCHAR(255) COMMENT '公司地址',
    logo VARCHAR(255) COMMENT '公司Logo',
    
    status VARCHAR(16) DEFAULT 'active' COMMENT '状态: active-正常 trial-试用 expired-已过期 inactive-停用',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    UNIQUE KEY uk_tenant_code (tenant_id, company_code),
    INDEX idx_tenant_id (tenant_id),
    INDEX idx_tenant_status (tenant_id, status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='公司表(租户表)';

-- 2. 自然人表
-- ============================================

-- 自然人表(跨公司的人员身份)
CREATE TABLE iam_person (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '自然人ID',
    
    -- 身份信息
    real_name VARCHAR(32) NOT NULL COMMENT '真实姓名',
    gender VARCHAR(8) COMMENT '性别: male-男 female-女 unknown-未知',
    birth_date DATE COMMENT '出生日期',
    
    -- 联系方式(自然人级别)
    mobile VARCHAR(16) COMMENT '手机号',
    email VARCHAR(64) COMMENT '邮箱',
    wechat VARCHAR(32) COMMENT '微信号',
    
    -- 账号信息(跨公司通用)
    password_hash VARCHAR(128) COMMENT '密码哈希(不存储盐值,盐值在应用层生成)',
    
    -- 其他信息
    avatar_url VARCHAR(255) COMMENT '头像URL',
    remark VARCHAR(500) COMMENT '备注',
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    INDEX idx_real_name (real_name),
    INDEX idx_mobile (mobile),
    INDEX idx_email (email),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='自然人表';

-- 3. 组织架构表
-- ============================================

-- 部门表
CREATE TABLE iam_department (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '部门ID',
    company_id BIGINT NOT NULL COMMENT '所属公司ID(租户ID)',
    parent_id BIGINT DEFAULT 0 COMMENT '父部门ID,0表示根部门',
    dept_code VARCHAR(32) NOT NULL COMMENT '部门编码',
    dept_name VARCHAR(64) NOT NULL COMMENT '部门名称',
    dept_path VARCHAR(512) COMMENT '部门路径: /1/2/3/',
    dept_level INT DEFAULT 1 COMMENT '部门层级',
    leader_id BIGINT COMMENT '部门负责人ID',
    sort_order INT DEFAULT 0 COMMENT '排序',
    status VARCHAR(16) DEFAULT 'active' COMMENT '状态: active-正常 inactive-停用',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    UNIQUE KEY uk_company_code (company_id, dept_code),
    INDEX idx_company_parent (company_id, parent_id),
    INDEX idx_company_path (company_id, dept_path(100)),
    INDEX idx_company_status (company_id, status),
    INDEX idx_leader (leader_id),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='部门表';

-- 4. 用户账号表
-- ============================================

-- 用户账号表(公司内的账号或平台管理员账号,一个自然人可在多个公司有账号)
CREATE TABLE iam_user (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
    person_id BIGINT NOT NULL COMMENT '自然人ID',
    company_id BIGINT NOT NULL DEFAULT 0 COMMENT '所属公司ID(租户ID), 0表示平台管理员账号',
    dept_id BIGINT COMMENT '主部门ID(冗余字段,方便查询,实际关联关系在iam_user_department表)',
    
    -- 账号信息
    username VARCHAR(32) NOT NULL COMMENT '用户名(公司用户:公司内唯一,平台管理员:全局唯一,需应用层保证)',
    
    -- 公司内信息(仅公司用户有效)
    employee_no VARCHAR(32) COMMENT '工号',
    position VARCHAR(64) COMMENT '职位',
    job_level VARCHAR(32) COMMENT '职级',
    entry_date DATE COMMENT '入职日期',
    
    -- 账号状态
    status VARCHAR(16) DEFAULT 'active' COMMENT '状态: active-正常 locked-锁定 disabled-禁用',
    user_type VARCHAR(16) DEFAULT 'normal' COMMENT '用户类型: normal-普通用户 company_admin-公司管理员 platform_admin-平台管理员(可管理所有公司)',
    last_login_at DATETIME(3) NULL COMMENT '最后登录时间',
    last_login_ip VARCHAR(45) COMMENT '最后登录IP(支持IPv6)',
    login_count INT DEFAULT 0 COMMENT '登录次数',
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    UNIQUE KEY uk_company_username (company_id, username),
    UNIQUE KEY uk_company_employee_no (company_id, employee_no),
    INDEX idx_person_id (person_id),
    INDEX idx_company_dept (company_id, dept_id),
    INDEX idx_company_status (company_id, status),
    INDEX idx_company_user_type (company_id, user_type),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户账号表';

-- 用户部门关联表(支持用户跨部门,每个用户只能有一个主部门,需应用层保证)
CREATE TABLE iam_user_department (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '关联ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    company_id BIGINT NOT NULL COMMENT '公司ID(租户ID,冗余)',
    dept_id BIGINT NOT NULL COMMENT '部门ID',
    dept_type VARCHAR(16) DEFAULT 'primary' COMMENT '部门类型: primary-主部门 secondary-其他部门',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    UNIQUE KEY uk_user_dept (user_id, dept_id),
    INDEX idx_user_id (user_id),
    INDEX idx_company_dept (company_id, dept_id),
    INDEX idx_user_type (user_id, dept_type),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户部门关联表';

-- 5. 权限管理表
-- ============================================

-- 角色表
CREATE TABLE iam_role (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '角色ID',
    company_id BIGINT NOT NULL COMMENT '所属公司ID(租户ID)',
    role_code VARCHAR(32) NOT NULL COMMENT '角色编码',
    role_name VARCHAR(64) NOT NULL COMMENT '角色名称',
    role_type VARCHAR(16) DEFAULT 'custom' COMMENT '角色类型: custom-自定义 system-系统内置',
    description VARCHAR(500) COMMENT '角色描述',
    data_scope VARCHAR(16) DEFAULT 'all' COMMENT '数据权限范围: all-全部 dept_and_sub-本部门及以下 dept-本部门 self-仅本人 custom-自定义',
    sort_order INT DEFAULT 0 COMMENT '排序',
    status VARCHAR(16) DEFAULT 'active' COMMENT '状态: active-正常 inactive-停用',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    UNIQUE KEY uk_company_code (company_id, role_code),
    INDEX idx_company_status (company_id, status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色表';

-- 用户角色关联表
CREATE TABLE iam_user_role (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '关联ID',
    company_id BIGINT NOT NULL COMMENT '公司ID(租户ID,冗余)',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    role_id BIGINT NOT NULL COMMENT '角色ID',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    UNIQUE KEY uk_user_role (company_id, user_id, role_id),
    INDEX idx_company_user (company_id, user_id),
    INDEX idx_company_role (company_id, role_id),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色关联表';

-- 菜单表
CREATE TABLE iam_menu (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '菜单ID',
    company_id BIGINT NOT NULL COMMENT '所属公司ID(租户ID)',
    parent_id BIGINT DEFAULT 0 COMMENT '父菜单ID',
    menu_code VARCHAR(32) NOT NULL COMMENT '菜单编码',
    menu_name VARCHAR(64) NOT NULL COMMENT '菜单名称',
    menu_type VARCHAR(16) DEFAULT 'directory' COMMENT '菜单类型: directory-目录 menu-菜单 button-按钮',
    
    -- 路由信息
    route_path VARCHAR(255) COMMENT '路由地址',
    component_path VARCHAR(255) COMMENT '组件路径',
    
    -- 权限标识
    permission VARCHAR(64) COMMENT '权限标识: sys:user:add',
    
    -- 显示信息
    icon VARCHAR(64) COMMENT '菜单图标',
    sort_order INT DEFAULT 0 COMMENT '排序',
    visibility VARCHAR(16) DEFAULT 'visible' COMMENT '可见性: visible-可见 hidden-隐藏',
    cache_type VARCHAR(16) DEFAULT 'disabled' COMMENT '缓存类型: enabled-启用 disabled-禁用',
    link_type VARCHAR(16) DEFAULT 'internal' COMMENT '链接类型: internal-内部链接 external-外部链接',
    
    status VARCHAR(16) DEFAULT 'active' COMMENT '状态: active-正常 inactive-停用',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    UNIQUE KEY uk_company_code (company_id, menu_code),
    INDEX idx_company_parent (company_id, parent_id),
    INDEX idx_company_status (company_id, status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单表';

-- 角色菜单关联表
CREATE TABLE iam_role_menu (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '关联ID',
    company_id BIGINT NOT NULL COMMENT '公司ID(租户ID,冗余)',
    role_id BIGINT NOT NULL COMMENT '角色ID',
    menu_id BIGINT NOT NULL COMMENT '菜单ID',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    UNIQUE KEY uk_role_menu (company_id, role_id, menu_id),
    INDEX idx_company_role (company_id, role_id),
    INDEX idx_company_menu (company_id, menu_id),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色菜单关联表';

-- 6. 审计日志表
-- ============================================

-- 操作日志表
CREATE TABLE iam_operation_log (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '日志ID',
    company_id BIGINT NOT NULL COMMENT '公司ID(租户ID)',
    user_id BIGINT COMMENT '操作人ID',
    username VARCHAR(32) COMMENT '操作人账号',
    
    module VARCHAR(32) COMMENT '操作模块',
    operation VARCHAR(16) COMMENT '操作类型: create/update/delete/query',
    method VARCHAR(16) COMMENT '请求方法: GET/POST/PUT/DELETE等',
    request_id VARCHAR(64) COMMENT '请求ID(用于追踪请求链路)',
    request_url VARCHAR(512) COMMENT '请求URL',
    request_params TEXT COMMENT '请求参数(JSON格式)',
    response_result TEXT COMMENT '返回结果(JSON格式)',
    
    ip_address VARCHAR(45) COMMENT 'IP地址(支持IPv6)',
    user_agent VARCHAR(512) COMMENT '用户代理',
    
    status VARCHAR(16) DEFAULT 'success' COMMENT '操作状态: success-成功 failed-失败',
    error_msg VARCHAR(1000) COMMENT '错误信息',
    execute_time INT COMMENT '执行时长(ms)',
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    INDEX idx_company_user (company_id, user_id),
    INDEX idx_company_created (company_id, created_at),
    INDEX idx_company_module (company_id, module),
    INDEX idx_request_id (request_id),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='操作日志表';

-- 登录日志表
CREATE TABLE iam_login_log (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '日志ID',
    company_id BIGINT NOT NULL COMMENT '公司ID(租户ID)',
    user_id BIGINT COMMENT '用户ID',
    username VARCHAR(32) COMMENT '用户名',
    
    login_type VARCHAR(16) COMMENT '登录类型: password/sms/wechat',
    login_status VARCHAR(16) COMMENT '登录状态: success-成功 failed-失败',
    login_message VARCHAR(128) COMMENT '登录消息',
    
    ip_address VARCHAR(45) COMMENT 'IP地址(支持IPv6)',
    location VARCHAR(128) COMMENT '登录地点',
    browser VARCHAR(64) COMMENT '浏览器',
    os VARCHAR(64) COMMENT '操作系统',
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3) NULL COMMENT '删除时间, NULL表示未删除',
    created_by BIGINT NOT NULL DEFAULT 0 COMMENT '创建人ID',
    updated_by BIGINT NOT NULL DEFAULT 0 COMMENT '更新人ID',
    deleted_by BIGINT NOT NULL DEFAULT 0 COMMENT '删除人ID',
    
    INDEX idx_company_user (company_id, user_id),
    INDEX idx_company_created (company_id, created_at),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='登录日志表';