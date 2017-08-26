<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link
	href="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1502527377034&di=bd3b84165bf98b08e52c2b7618fdb853&imgtype=0&src=http%3A%2F%2Fec4.images-amazon.com%2Fimages%2FI%2F61QFHDUcvQL._SL1000_.jpg"
	rel="shortcut icon" type="image/x-icon" />
<meta name="description" content="录入授权信息">
<meta name="author" content="JerryChu">
<!-- Bootstrap core CSS -->
<link href="//cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css"
	rel="stylesheet">
<title>录入授权信息</title>
</head>
<body>
	<nav class="navbar navbar-inverse navbar-fixed-top">
		<div class="container">
			<div class="navbar-header">
				<a class="navbar-brand" href="/auth">Personal Authorization
					Information</a>
			</div>
		</div>
	</nav>

	<div class="jumbotron">
		<div class="container">
		    <h3>...</h3>
		</div>
	</div>

	<div class="container">
		<div class="row">
			<div class="col-md-2"></div>

			<div class="col-md-8">
				<form class="form-horizontal" id="userForm" method="POST"
					action="auth/user/save">

					<div class="form-group">
						<label for="Src" class="col-sm-2 control-label">Src</label>
						<div class="col-sm-3">
							<select class="form-control" id="src" name="src">
                              <option value="wish">Wish</option>
                              <option value="ebay">Ebay</option>
                              <option value="th">Th</option>
                              <option value="znu">Znu</option>
                            </select>
						</div>
					</div>

					<div class="form-group">
						<label for="UserCode" class="col-sm-2 control-label">UserCode</label>
						<div class="col-sm-3">
							<input type="text" class="form-control" id="user_code"
								name="user_code" placeholder="UserCode">
						</div>
					</div>

					<div class="form-group">
						<label for="UserNick" class="col-sm-2 control-label">UserNick</label>
						<div class="col-sm-4">
							<input type="text" class="form-control" id="user_nick"
								name="user_nick" placeholder="UserNick">
						</div>
					</div>

					<div class="form-group">
						<label for="Mobile" class="col-sm-2 control-label">Mobile</label>
						<div class="col-sm-5">
							<input type="phone" class="form-control" id="mobile"
								name="mobile" placeholder="Mobile">
						</div>
					</div>

					<div class="form-group">
						<label for="Email" class="col-sm-2 control-label">Email</label>
						<div class="col-sm-5">
							<input type="email" class="form-control" id="email"
							    name="email" placeholder="Email">
						</div>
					</div>

					<div class="form-group">
						<label for="Secret" class="col-sm-2 control-label">Secret</label>
						<div class="col-sm-6">
							<input type="text" class="form-control" id="secret" name="secret"
								placeholder="Secret">
						</div>
					</div>

					<div class="form-group">
						<label for="Info" class="col-sm-2 control-label">Info</label>
						<div class="col-sm-10">
							<textarea class="form-control" rows="5" id="info" name="info"></textarea>
						</div>
					</div>

					<div class="form-group">
						<div class="col-sm-offset-6 col-sm-4">
							<input type="submit" class="btn btn-primary" name="submit" onclick="check()"
								value="添加信息">
						</div>
					</div>
				</form>
			</div>

			<div class="col-md-2"></div>

		</div>

		<hr>

		<footer>
			<nav class="navbar navbar-inverse ">
				<div class="container">
					<div class="navbar-header">
						<a class="navbar-brand" href="#">Welcome to go
							programming...</a>
					</div>
				</div>
			</nav>
		</footer>
	</div>
	<!-- /container -->
	<script type="text/javascript"
		src="http://code.jquery.com/jquery-latest.js"></script>
	<script type="text/javascript">
	    function check() {
	        $("#src").blur(function(){
	            var src = $("#src").val();
	            if (src == "") {
                    alert("src不能为空，请填写！");
                    return;
	            }
            });
            $("#email").blur(function(){
                var email = $("#email").val();
                if (email == "") {
                    alert("email不能为空，请填写！");
                    return;
                }
            });
	    }
	    $(document).ready(function(){
            check();
	    });
    </script>
    <script src="//cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</body>
</html>
