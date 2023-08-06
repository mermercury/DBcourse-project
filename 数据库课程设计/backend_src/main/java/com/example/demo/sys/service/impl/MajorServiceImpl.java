package com.example.demo.sys.service.impl;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.dao.mapper.ClassMapper;
import com.example.demo.sys.dao.mapper.DepartmentMapper;
import com.example.demo.sys.dao.mapper.StudentMapper;
import com.example.demo.sys.entity.Major;
import com.example.demo.sys.dao.mapper.MajorMapper;
import com.example.demo.sys.service.IMajorService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.HashMap;
import java.util.List;
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
public class MajorServiceImpl extends ServiceImpl<MajorMapper, Major> implements IMajorService {
    @Autowired
    private MajorMapper majorMapper;
    @Autowired
    private DepartmentMapper departmentMapper;
    @Autowired
    private StudentMapper studentMapper;
    @Autowired
    private ClassMapper classMapper;
    
    @Override
    public Result getMajorPageCount(String departmentName, String majorName, int pageSize) {
        int count = 0;
        if ((departmentName == null || departmentName.equals("")) && (majorName == null || majorName.equals(""))) {
            // 查询所有专业
            count = majorMapper.getMajorCount();
        } else if (departmentName == null || departmentName.equals("")) {
            count = majorMapper.getMajorCountByMajorName(majorName);
        } else if ((majorName == null || majorName.equals(""))) {
            count = majorMapper.getMajorCountByDepartmentName(departmentName);
        } else {
            count = majorMapper.getMajorCountByBoth(majorName, departmentName);
        }
        Map<String, Object> data = new HashMap<>();
        count  = (int) Math.ceil(count * 1.0 / pageSize);
        data.put("count", count);
        return Result.success(data);
    }

    @Override
    public Result getMajorByPage(String departmentName, String majorName, int offset, int limit) {
        List<Map<String, Object>> data = null;
        if ((departmentName == null || departmentName.equals("")) && (majorName == null || majorName.equals(""))) {
            // 查询所有课程
            data = majorMapper.getMajorByPage(offset, limit);
        } else if (departmentName == null || departmentName.equals("")) {
            data = majorMapper.getMajorByPageByMajorName(majorName, offset, limit);
        } else if ((majorName == null || majorName.equals(""))) {
            data = majorMapper.getMajorByPageByDepartmentName(departmentName, offset, limit);
        } else {
            data = majorMapper.getMajorByPageByBoth(departmentName, majorName, offset, limit);
        }
        return Result.success(data);
    }

    @Override
    public Result getAMajor(int id) {
        Map<String, Object> data = majorMapper.getMajorById(id);
        return Result.success(data);
    }

    @Override
    public Result updateMajorName(int majorId, String majorNewName) {
        majorMapper.updateMajorName(majorId, majorNewName);
        return Result.success();
    }

    @Override
    public Result createMajor(String majorName, String departmentName) {
        Major major = new Major();
        major.setMajorName(majorName);
        int departmentId = departmentMapper.getDepartmentIdByName(departmentName);
        major.setDepartmentId(departmentId);
        majorMapper.insert(major);
        departmentMapper.increaseMajorCount(departmentId);
        return Result.success();
    }

    @Transactional
    @Override
    public Result deleteAMajor(int id) {
        Major major = majorMapper.selectById(id);

        // 1. 删除该专业所有学生
        studentMapper.deleteStudentsByMajorId(id);


        // 2. 对应department的major_count减1
        int departmentId = major.getDepartmentId();
        departmentMapper.decreaseMajorCount(departmentId);

        // 3. 删除该专业所有班级
        classMapper.deleteClassesByMajorId(id);

        // 4. 删除该专业记录
        majorMapper.deleteById(id);

        return Result.success();
    }

}
