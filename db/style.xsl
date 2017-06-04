<?xml version="1.0" encoding="utf8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:s="http://sqlfairy.sourceforge.net/sqlfairy.xml">
	<xsl:output method="html" encoding="utf8" doctype-public="-//W3C//DTD XHTML 1.0 Transitional//EN"/>

	<xsl:template match="database">
		<html lang="ja">
			<head>
				<meta charset="utf-8"/>
				<title>テーブル定義</title>
				<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css"/>
				<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap-theme.min.css"/>
				<style>
					table.htable {
						margin: 3em auto 1em auto !important;
					}
					table.htable th {
						border-left: 10px solid #e5e5e5 !important;
					}
					footer {
						border-top: 1px solid #e5e5e5;
						padding: 0.5em;
					}
				</style>
			</head>

			<body>
				<div class="container">
					<h1 class="page-header">テーブル定義</h1>
                    <ul class="list-group">
                        <li class="list-group-item list-group-item-info">テーブル</li>
                        <xsl:apply-templates select="table_structure" mode="list"/>
                    </ul>
					<xsl:apply-templates select="table_structure"/>
				</div>
				<footer class="text-center">
					your company
				</footer>
			</body>
		</html>
	</xsl:template>

    <xsl:template match="table_structure" mode="list">
        <li class="list-group-item">
            <a>
                <xsl:attribute name="href">
                    #<xsl:value-of select="@name" />
                </xsl:attribute>
                <xsl:value-of select="@name" />（<xsl:value-of select="options/@Comment"/>）
            </a>
        </li>
	</xsl:template>

	<xsl:template match="table_structure">
		<table class="table table-bordered htable">
            <xsl:attribute name="id">
                <xsl:value-of select="@name" />
            </xsl:attribute>
			<tbody>
				<tr class="active">
					<th>テーブル名</th>
					<td>
						<xsl:value-of select="options/@Comment"/>
					</td>
				</tr>
				<tr class="active">
					<th>スキーマ</th>
					<td><xsl:value-of select="@name"/></td>
				</tr>
			</tbody>
		</table>
		<table class="table table-condensed">
			<thead>
				<tr>
					<th class="text-right">#</th>
					<th>論理名</th>
					<th>物理名</th>
					<th>型</th>
					<th>NULL</th>
					<th>デフォルト値</th>
					<th>主キー</th>
					<th>ユニーク</th>
				</tr>
			</thead>
			<tbody>
				<xsl:apply-templates select="field"/>
			</tbody>
		</table>
	</xsl:template>

	<xsl:template match="field">
		<tr>
			<td class="text-right"><xsl:value-of select="position()"/></td>
			<td><xsl:value-of select="@Comment"/></td>
			<td><xsl:value-of select="@Field"/></td>
			<td><xsl:value-of select="@Type"/></td>
			<td><xsl:if test="@Null='YES'"><span class="glyphicon glyphicon-ok"></span></xsl:if></td>
			<td><xsl:value-of select="@Default"/></td>
			<td><xsl:if test="@Key='PRI'"><span class="glyphicon glyphicon-ok"></span></xsl:if></td>
			<td><xsl:if test="@Key='UNI'"><span class="glyphicon glyphicon-ok"></span></xsl:if></td>
		</tr>
	</xsl:template>
</xsl:stylesheet>
