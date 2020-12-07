<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN
">
<html>
<head>
<title>Manifesto for Agile Software Development
</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<META name="description" content="We are uncovering better ways of developing software
by doing it and helping others do it. These are our
values and principles.
">
<META name="keywords" content="agilemanifesto agilealliance alliance manifesto,
agile programming, agile development,
extreme programming, XP, software project management,
iterative development, collaborative development,
software engineering best practices, software development,
best practices
">
</head>
<body background="background.jpg" style="background-size: contain;">
<center>
<br><br><br><br>
<?php

if (isset($_GET["lang"])){
	$lang = $_GET["lang"];

	if ($lang === "index.php"){
		die("Too many redirects");
	}

	$bad = array("bin", "boot", "dev", "etc", "home", "lib", "lib64", "media", "mnt", "opt", "proc", "root", "run", "run.sh", "sbin", "srv", "sys", "tmp", "usr", "var", "x.sql", "~", "..");

	foreach ($bad as $b){
		if (strpos($lang, $b) !== false){
			die("Hacker detected! Go away!");
		}
	}

	include($lang);
}else{
	include("en.php");
}

?>

<font size="+2">
<a href="#">Twelve Principles of Agile Software</a><br><br>
<a href="#">View Signatories</a><br><br>
<a href="#">About the Authors</a><br>
<a href="#">About the Manifesto</a><br>

</font>
<br><br>

<font size="+2">
<a href="index.php?lang=en.php">English</a><br>
<a href="index.php?lang=es.php">Español</a><br>
<a href="index.php?lang=fr.php">Français</a><br>
<a href="index.php?lang=sw.php">Swahili</a><br>

</font>
<br><br>

<font size=1 color=gray>
site design and artwork &copy; 2001, Ward Cunningham

</font>

</center>
</body>
</html>

