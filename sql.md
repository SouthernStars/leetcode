-- 选择每个人的最大分数

SELECT s.student_id, s.course_id, s.grade, t.max_grade FROM sc as s LEFT JOIN
(
SELECT student_id, max(grade) as max_grade from sc GROUP BY student_id
)
AS t ON s.student_id = t.student_id WHERE s.grade = t.max_grade ORDER BY s.student_id ASC, s.course_id ASC;

