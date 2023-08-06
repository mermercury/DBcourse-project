package com.example.demo.sys.service.impl;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.Class;
import com.example.demo.sys.dao.mapper.ClassMapper;
import com.example.demo.sys.service.IClassService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Map;

/**
 * <p>
 *  服务实现类
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
@Service
public class ClassServiceImpl extends ServiceImpl<ClassMapper, Class> implements IClassService {
    @Autowired
    private ClassMapper classMapper;

    @Override
    public Result getClassInfo(int classId) {
        Map<String, Object> data = classMapper.getAClassById(classId);
        return Result.success(data);
    }

    @Override
    public Result getClassPageCount(String majorName, String departmentName, String studentName, int pageSize) {


        return null;
    }
}
