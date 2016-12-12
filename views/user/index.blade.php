@extends('layouts.common')
@section('content')
<div class=row>
    <div class=col-md-12>
        <div class="panel mb25">
            <div class=panel-heading>用户资料</div>
            <div class=panel-body>
                <table class="table mb0">
                  <thead>
                    <tr>
                      <th>＃</th>
                      <th>用户名</th>
                      <th>手机</th>
                      <th>性别</th>
                      <th>注册类型</th>
                      <th>third_id</th>
                      <th>交易账号</th>
                      <th>注册日期</th>
                      <th>用户状态</th>
                      <th><a>操作</a></th>
                    </tr>
                  </thead>
                  <tbody>
                    <?php $status_arr = ['1'=>'正常','0'=>'禁止登录'];?>
                    <?php foreach ($user as $k=>$v) {
    ?>
                    <tr>
                      <td>{{$v->id}}</td>
                      <td>{{$v->name}}</td>
                      <td>{{$v->telephone}}</td>
                      <td>{{$v->sex}}</td>
                      <td>{{$v->type}}</td>
                      <td>{{$v->third_id}}</td>
                      <td>{{$v->trade}}</td>
                      <td>{{$v->created_at}}</td>
                      <td>{{$status_arr[$v->status]}}</td>
                      <td><a href="#">详情</a> <a class="alink" href="/admin/user/delete?id={{$v->id}}">删除</a></td>
                    </tr>
                    <?php

}?>
                  </tbody>

                </table>
            </div>
        </div>
</div>
</div>
<?php echo $user->render(); ?>
<script type="text/javascript">
$(".alink").bind('click',function(e){
  var link = new Link(e);
  link.tip();
})
</script>
@stop
