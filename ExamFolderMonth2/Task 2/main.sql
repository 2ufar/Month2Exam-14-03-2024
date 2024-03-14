SELECT
    e.name AS Employee_Name,
    d.name AS Job_Title,
    COUNT(s.id) AS Number_Of_Times_Received_Salary
FROM
    employees e
JOIN
    departments d ON e.department_id = d.id
LEFT JOIN
    salaries s ON e.id = s.employee_id
GROUP BY
    e.name, d.name
ORDER BY
    e.name;
