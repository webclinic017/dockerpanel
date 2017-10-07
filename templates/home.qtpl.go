// This file is automatically generated by qtc from "home.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/home.qtpl:1
package templates

//line templates/home.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/home.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/home.qtpl:1
func StreamHomeTemplate(qw422016 *qt422016.Writer) {
	//line templates/home.qtpl:1
	qw422016.N().S(`

<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Panel Docker</title>

    <!-- Bootstrap Core CSS -->
    <link href="../vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">

    <!-- MetisMenu CSS -->
    <link href="../vendor/metisMenu/metisMenu.min.css" rel="stylesheet">

    <!-- Custom CSS -->
    <link href="../dist/css/sb-admin-2.css" rel="stylesheet">

    <!-- Morris Charts CSS -->
    <link href="../vendor/morrisjs/morris.css" rel="stylesheet">

    <!-- Custom Fonts -->
    <link href="../vendor/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
    <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

</head>

<body>

<div id="wrapper">

    <!-- Navigation -->
    <nav class="navbar navbar-default navbar-static-top" role="navigation" style="margin-bottom: 0">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="index.html">Docker Panel</a>
        </div>
        <!-- /.navbar-header -->

        <ul class="nav navbar-top-links navbar-right">
            <!-- /.dropdown -->
            <li class="dropdown">
                <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                    <i class="fa fa-user fa-fw"></i> <i class="fa fa-caret-down"></i>
                </a>
                <ul class="dropdown-menu dropdown-user">
                    <li><a href="#"><i class="fa fa-user fa-fw"></i> User Profile</a>
                    </li>
                    <li class="divider"></li>
                    <li><a href="/logout"><i class="fa fa-sign-out fa-fw"></i> Logout</a>
                    </li>
                </ul>
                <!-- /.dropdown-user -->
            </li>
            <!-- /.dropdown -->
        </ul>
        <!-- /.navbar-top-links -->

        <div class="navbar-default sidebar" role="navigation">
            <div class="sidebar-nav navbar-collapse">
                <ul class="nav" id="side-menu">
                    <li>
                        <a href="/"><i class="fa fa-dashboard fa-fw"></i> Dashboard</a>
                    </li>
                    <li>
                        <a href="/web"><i class="fa fa-server fa-fw"></i>Sitios Web</a>
                    </li>
                    <li>
                        <a href="/dns"><i class="fa fa-cloud fa-fw"></i>DNS</a>
                    </li>
                    <li>
                        <a href="/db"><i class="fa fa-database fa-fw"></i>Base de Datos</a>
                    </li>
                    <li>
                        <a href="/mail"><i class="fa fa-at fa-fw"></i>E-Mail</a>
                    </li>
                    <li>
                        <a href="/ftp"><i class="fa fa-file-archive-o fa-fw"></i>FTP</a>
                    </li>
                </ul>
            </div>
            <!-- /.sidebar-collapse -->
        </div>
        <!-- /.navbar-static-side -->
    </nav>

    <div id="page-wrapper">




    </div>
    <!-- /#page-wrapper -->

</div>
<!-- /#wrapper -->

<!-- jQuery -->
<script src="../vendor/jquery/jquery.min.js"></script>

<!-- Bootstrap Core JavaScript -->
<script src="../vendor/bootstrap/js/bootstrap.min.js"></script>

<!-- Metis Menu Plugin JavaScript -->
<script src="../vendor/metisMenu/metisMenu.min.js"></script>

<!-- Morris Charts JavaScript -->
<!--<script src="../vendor/raphael/raphael.min.js"></script>-->
<!--<script src="../vendor/morrisjs/morris.min.js"></script>-->
<!--<script src="../data/morris-data.js"></script>-->

<!-- Custom Theme JavaScript -->
<script src="../dist/js/sb-admin-2.js"></script>

</body>

</html>

`)
//line templates/home.qtpl:136
}

//line templates/home.qtpl:136
func WriteHomeTemplate(qq422016 qtio422016.Writer) {
	//line templates/home.qtpl:136
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/home.qtpl:136
	StreamHomeTemplate(qw422016)
	//line templates/home.qtpl:136
	qt422016.ReleaseWriter(qw422016)
//line templates/home.qtpl:136
}

//line templates/home.qtpl:136
func HomeTemplate() string {
	//line templates/home.qtpl:136
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/home.qtpl:136
	WriteHomeTemplate(qb422016)
	//line templates/home.qtpl:136
	qs422016 := string(qb422016.B)
	//line templates/home.qtpl:136
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/home.qtpl:136
	return qs422016
//line templates/home.qtpl:136
}
