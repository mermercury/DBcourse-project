package com.example.demo.sys.service;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.Class;
import com.baomidou.mybatisplus.extension.service.IService;

/**
 * <p>
 *  服务类
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
public interface IClassService extends IService<Class> {

    Result getClassInfo(int classId);

    Result getClassPageCount(String majorName, String departmentName, String studentName, int pageSize);

    Result getClassByPage(String majorName, String departmentName, String studentName, int limit);
}
