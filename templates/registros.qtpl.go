// This file is automatically generated by qtc from "registros.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/registros.qtpl:1
package templates

//line templates/registros.qtpl:1
import "github.com/mkreder/dockerpanel/model"

//line templates/registros.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/registros.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/registros.qtpl:2
func StreamRegistroTemplate(qw422016 *qt422016.Writer, registros []model.Registro, zonaid string, error string) {
	//line templates/registros.qtpl:2
	qw422016.N().S(`

<!DOCTYPE html>
<html lang="en">
<script type="text/javascript">
    function validateForm() {
        var x = document.getElementById("nombre").value;
        if (x == "" ) {
            alert("Se debe completar el nombre");
            return false;
        }

        var y = document.getElementById("valor").value;
        if (y == "" ) {
            alert("Se debe completar el valor");
            return false;
        }

        var z = document.getElementById("prioridad").value;
        if (z == "" ) {
            alert("Se debe completar la prioridad");
            return false;
        }

        if (/^[0-9]*$/.test(z)) {
        } else {
            alert("Prioridad debe ser un valor numerico");
            return false;
        }


    }

    function showDiv() {
        var x = document.getElementById('form');
        x.style.display = 'block';
    }


    function hideDiv() {
        var x = document.getElementById('form');
        x.style.display = 'none';
        document.getElementById("tipo").options[0].selected = true;
        document.getElementById("nombre").value = "";
        document.getElementById("valor").value = "";
        document.getElementById("prioridad").value = "";
        document.getElementById("id").value = "";
    }

    function modifyRegistro(id,tipo,nombre,valor,prioridad) {
        var x = document.getElementById('form');
        document.getElementById("id").value = id;
        var seltipo = document.getElementById("tipo");


        for (var i = 0; i < seltipo.options.length; i++) {
            if (seltipo.options[i].text== tipo) {
                seltipo.options[i].selected = true;
            }
        }
        document.getElementById("nombre").value = nombre;
        document.getElementById("valor").value = valor;
        document.getElementById("prioridad").value = prioridad;

        x.style.display = 'block';
    }

</script>
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

`)
	//line templates/registros.qtpl:106
	if len(error) > 0 {
		//line templates/registros.qtpl:106
		qw422016.N().S(`
<script type="text/javascript">
    window.alert("`)
		//line templates/registros.qtpl:108
		qw422016.E().S(error)
		//line templates/registros.qtpl:108
		qw422016.N().S(`")
</script>
`)
		//line templates/registros.qtpl:110
	}
	//line templates/registros.qtpl:110
	qw422016.N().S(`

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
                    <li><a href="/profile"><i class="fa fa-user fa-fw"></i> Configuración</a>
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
        <br>
        <div class="row">
            <div class="col-lg-12">
                <div class="panel panel-default">
                    <div class="panel-heading clearfix">
                        <h4 class="panel-title pull-left" style="padding-top: 7.5px;">Registros</h4>
                        <button type="button" class="btn pull-right btn-primary btn-sm" onclick="showDiv()">Agregar</button>
                    </div>
                    <!-- /.panel-heading -->
                    <div class="panel-body">
                        <table width="100%" class="table table-striped table-bordered table-hover" id="dataTables-example">
                            <thead>
                            <tr>
                                <th>Tipo</th>
                                <th>Nombre</th>
                                <th>Valor</th>
                                <th>Prioridad</th>
                                <th>Acciones</th>
                            </tr>
                            </thead>
                            <tbody>
                            `)
	//line templates/registros.qtpl:196
	for _, registro := range registros {
		//line templates/registros.qtpl:196
		qw422016.N().S(`
                            <tr class="odd gradeX">
                                <td> `)
		//line templates/registros.qtpl:198
		qw422016.E().S(registro.Tipo)
		//line templates/registros.qtpl:198
		qw422016.N().S(` </td>
                                <td> `)
		//line templates/registros.qtpl:199
		qw422016.E().S(registro.Nombre)
		//line templates/registros.qtpl:199
		qw422016.N().S(` </td>
                                <td> `)
		//line templates/registros.qtpl:200
		qw422016.E().S(registro.Valor)
		//line templates/registros.qtpl:200
		qw422016.N().S(` </td>
                                <td> `)
		//line templates/registros.qtpl:201
		qw422016.E().S(registro.Prioridad)
		//line templates/registros.qtpl:201
		qw422016.N().S(` </td>

                                <td class="center">
                                    <button type="button" class="btn btn-xs btn-primary" onclick='modifyRegistro(`)
		//line templates/registros.qtpl:204
		qw422016.N().D(int(registro.ID))
		//line templates/registros.qtpl:204
		qw422016.N().S(`, "`)
		//line templates/registros.qtpl:204
		qw422016.E().S(registro.Tipo)
		//line templates/registros.qtpl:204
		qw422016.N().S(`",  "`)
		//line templates/registros.qtpl:204
		qw422016.E().S(registro.Nombre)
		//line templates/registros.qtpl:204
		qw422016.N().S(`", "`)
		//line templates/registros.qtpl:204
		qw422016.E().S(registro.Valor)
		//line templates/registros.qtpl:204
		qw422016.N().S(`",  "`)
		//line templates/registros.qtpl:204
		qw422016.E().S(registro.Prioridad)
		//line templates/registros.qtpl:204
		qw422016.N().S(`" )' ><i class="fa fa-list"></i></button>
                                    <button class="btn btn-xs btn-danger" onclick="location.href='removeRegistro?id=`)
		//line templates/registros.qtpl:205
		qw422016.N().D(int(registro.ID))
		//line templates/registros.qtpl:205
		qw422016.N().S(`&zonaid=`)
		//line templates/registros.qtpl:205
		qw422016.E().S(zonaid)
		//line templates/registros.qtpl:205
		qw422016.N().S(`';"><i class="fa fa-trash-o"></i></button>
                                </td>
                            </tr>
                            `)
		//line templates/registros.qtpl:208
	}
	//line templates/registros.qtpl:208
	qw422016.N().S(`
                            </tbody>
                        </table>
                    </div>
                    <!-- /.panel-body -->
                </div>
                <!-- /.panel -->
            </div>
            <!-- /.col-lg-12 -->
        </div>


        <div id="form" class="col-lg-6" hidden="true" >
            <div class="panel panel-default">
                <div class="panel-heading">
                    Nuevo Sitio
                </div>
                <div class="panel-body">
                    <form id="addregistro" action="/registros" onsubmit="return validateForm()" role=form method="post">
                        <input id="id" name="id" id="id" hidden="true" >
                        <input id="zonaid" name="zonaid" id="zonaid" value="`)
	//line templates/registros.qtpl:228
	qw422016.E().S(zonaid)
	//line templates/registros.qtpl:228
	qw422016.N().S(`" hidden="true" >
                        Tipo
                        <select  class="form-control"  id="tipo" name="tipo">
                            <option value="NS">NS</option>
                            <option value="A">A</option>
                            <option value="AAAA">AAAA</option>
                            <option value="CNAME">CNAME</option>
                            <option value="MX">MX</option>
                            <option value="PTR">PTR</option>
                            <option value="TXT">TXT</option>
                            <option value="SRV">SRV</option>
                            <option value="DS">DS</option>
                        </select>
                        <br>
                        Nombre
                        <input class="form-control" name="nombre" id="nombre">
                        <br>
                        Valor
                        <input class="form-control" name="valor" id="valor">
                        <br>
                        Prioridad
                        <input class="form-control" name="prioridad" id="prioridad">
                        <br>
                        <button type="submit" class="btn btn-default">Guardar</button>
                        <button type="button" class="btn btn-default" onclick="hideDiv()">Cancelar</button>
                    </form>
                </div>
            </div>

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
//line templates/registros.qtpl:284
}

//line templates/registros.qtpl:284
func WriteRegistroTemplate(qq422016 qtio422016.Writer, registros []model.Registro, zonaid string, error string) {
	//line templates/registros.qtpl:284
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/registros.qtpl:284
	StreamRegistroTemplate(qw422016, registros, zonaid, error)
	//line templates/registros.qtpl:284
	qt422016.ReleaseWriter(qw422016)
//line templates/registros.qtpl:284
}

//line templates/registros.qtpl:284
func RegistroTemplate(registros []model.Registro, zonaid string, error string) string {
	//line templates/registros.qtpl:284
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/registros.qtpl:284
	WriteRegistroTemplate(qb422016, registros, zonaid, error)
	//line templates/registros.qtpl:284
	qs422016 := string(qb422016.B)
	//line templates/registros.qtpl:284
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/registros.qtpl:284
	return qs422016
//line templates/registros.qtpl:284
}
