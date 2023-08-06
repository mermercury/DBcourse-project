package com.example.demo.sys.service.impl;

import com.example.demo.sys.entity.Course;
import com.example.demo.sys.dao.mapper.CourseMapper;
import com.example.demo.sys.service.ICourseService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * <p>
 *  服务实现类
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
@Service
public class CourseServiceImpl extends ServiceImpl<CourseMapper, Course> implements ICourseService {

}
