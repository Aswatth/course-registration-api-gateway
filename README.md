This is the API-Gateway for all micro-services of course-registration-system, acting as an entry point into the system by handling authentication and authorization using <a href="https://jwt.io/">JWT tokens</a> and managing communication across micro-services.

<h1>Micro services</h1>
<table>
  <tr>
    <th>Micro-serivce</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><a href="https://github.com/Aswatth/course-registration-profile-service">Profile service</a></td>
    <td>Mangaes profiles for all users</td>
  </tr>
  <tr>
    <td><a href="https://github.com/Aswatth/course-registration-course-service">Course service</a></td>
    <td>Manages courses which can be offered by professors</td>
  </tr>
  <tr>
    <td><a href="https://github.com/Aswatth/course-registration-registration-service">Registration service</a></td>
    <td>Manages offered courses for professors and registered courses for students</td>
  </tr>
</table>

<h1>API Endpoints</h1>
<h3>Login</h3>
<table>
  <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>POST</td>
    <td>/login</td>
    <td>Login using email_id and password</td>
  </tr>
</table>
<hr>

<h3>Admin profile endpoints</h3>
<table>
  <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/admin/password/:email_id</td>
    <td>Update admin profile password</td>
  </tr>
  <tr>
    <td>POST</td>
    <td>/admin/students</td>
    <td>Create new student profile</td>
  </tr>
  <tr>
    <td>GET</td>
    <td>/admin/students</td>
    <td>Get all student profile</td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/admin/students/:email_id</td>
    <td>Update a student profile using email_id</td>
  </tr> 
  <tr>
    <td>DELETE</td>
    <td>/admin/students/:email_id</td>
    <td>Delete a student profile using email_id</td>
  </tr>
  <tr>
    <td>POST</td>
    <td>/admin/professors</td>
    <td>Create new professor profile</td>
  </tr>
  <tr>
    <td>GET</td>
    <td>/admin/professors</td>
    <td>Get all professor profile</td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/admin/professors/:email_id</td>
    <td>Update a professor profile using email_id</td>
  </tr> 
  <tr>
    <td>DELETE</td>
    <td>/admin/professors/:email_id</td>
    <td>Delete a professor profile using email_id</td>
  </tr>
</table>
<hr>
<h3>Admin course endpoints</h3>
<table>
  <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>POST</td>
    <td>/admin/courses</td>
    <td>Add new course</td>
  </tr>
  <tr>
    <td>GET</td>
    <td>/admin/courses</td>
    <td>Fetch all courses</td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/admin/courses/:course_id</td>
    <td>Update a course using course_id</td>
  </tr>
  
  <tr>
    <td>DELETE</td>
    <td>/admin/courses/:course_id</td>
    <td>Delete a course using course_id</td>
  </tr>
</table>
<hr>
<h3>Professor endpoints</h3>
<table>
   <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>GET</td>
    <td>/professors/:email_id</td>
    <td>Fetch a professor's profile using email_id</td>
  </tr>
   <tr>
    <td>PUT</td>
    <td>/professors/password/:email_id</td>
    <td>Update professor profile password</td>
  </tr>
  <tr>
    <td>GET</td>
    <td>/professors/courses</td>
    <td>Fetch all courses</td>
  </tr>
  <tr>
    <td>POST</td>
    <td>/professors/offered_course</td>
    <td>Offer a new course from available courses</td>
  </tr>
   <tr>
    <td>GET</td>
    <td>/professors/offered_course?crn=? <br>/professors/offered_course?email_id=? </td>
    <td>Fetch all offered courses by CRN or email id</td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/professors/offered_course/:crn</td>
    <td>Update an offered course</td>
  </tr>
  
  <tr>
    <td>DELETE</td>
    <td>/professors/offered_course/:crn</td>
    <td>Delete an offered course</td>
  </tr>
</table>
<hr>
<h3>Professor endpoints</h3>
<table>
   <tr>
    <th>Action</th>
    <th>URL</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>GET</td>
    <td>/students/:email_id</td>
    <td>Fetch a student's profile using email_id</td>
  </tr>
   <tr>
    <td>PUT</td>
    <td>/students/password/:email_id</td>
    <td>Update professor profile password</td>
  </tr>
<tr>
  <td>GET</td>
  <td>/students/offered_courses</td>
  <td>Fetch all offered courses</td>
</tr>
  <tr>
    <td>POST</td>
    <td>/students/register_course</td>
    <td>Register for an offered course</td>
  </tr>
  <tr>
    <td>GET</td>
    <td>/students/register_course?email_id=?</td>
    <td>Get all registered courses for a student</td>
  </tr>
  <tr>
    <td>PUT</td>
    <td>/students/register_course/:email_id</td>
    <td>Update registered courses for a student</td>
  </tr>
  <tr>
    <td>DELETE</td>
    <td>/students/register_course/:email_id</td>
    <td>Delete registered courses for a student</td>
  </tr>
  
</table>
