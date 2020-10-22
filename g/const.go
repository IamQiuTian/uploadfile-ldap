package g


const LoginTpl = `
<!DOCTYPE html>
<html>
<head>
	<title>Rise File Uploads</title>
</head>

<body bgcolor='#3284D6'>
	<h1>Rise File Uploads</h1>
	Login with your LDAP credentials.
	<p />
	<form method="POST" action="/login">
		Please fill in your LDAP username
		<input type="text" name="username" placeholder="username">
		<br>
		Please fill in your LDAP password
		<input type="password" name="password" placeholder="password">
		<p />
		<input type="submit" value="Login">
	</form>
</body>
</html>`

const UploadTpl = `
<!DOCTYPE html>
<html>
<head>
	<title>Upload File</title>
</head>

<body bgcolor='#3284D6'>
	<form method="POST" enctype="multipart/form-data" action="/upload">
	  <b>Select a file to upload</b><p />
	  <input type="file" name="fileupload" value="fileupload" id="fileupload">
		<p />
	  <input type="submit" value="Upload File">
	</form>
</body>
</html>`

const DoneTpl = `
<!DOCTYPE html>
<html>
<head>
	<title>File Uploaded</title>
	<meta http-equiv="refresh" content="3;url=/file/{{.Path}}">
</head>
<body bgcolor='#3284D6'>
<b>File Uploaded</b>
<p />
You will be redirected to the file chooser in 3 seconds.
</body>
</html>`