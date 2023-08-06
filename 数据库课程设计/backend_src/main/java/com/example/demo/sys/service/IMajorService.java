package com.example.demo.sys.service;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.Major;
import com.baomidou.mybatisplus.extension.service.IService;

/**
 * <p>
 *  服务类
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
public interface IMajorService extends IService<Major> {

    Result getMajorPageCount(String departmentName, String majorName, int pageSize);

    Result getMajorByPage(String departmentName, String majorName, int offset, int limit);

    Result getAMajor(int id);

    Result updateMajorName(int majorId, String majorNewName);

    Result createMajor(String majorName, String departmentName);

    Result deleteAMajor(int id);
}
