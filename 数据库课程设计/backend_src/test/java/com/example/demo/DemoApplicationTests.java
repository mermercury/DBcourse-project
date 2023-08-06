package com.example.demo;

import com.example.demo.sys.entity.Department;
import com.example.demo.sys.dao.mapper.DepartmentMapper;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import javax.annotation.Resource;
import java.util.List;

@SpringBootTest
class DemoApplicationTests {

	@Resource
	private DepartmentMapper departmentMapper;

	@Test
	void testMapper() {
		List<Department> departments = departmentMapper.selectList(null);
		departments.forEach(System.out::println);
	}

}
