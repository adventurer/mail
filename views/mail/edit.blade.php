@extends('layouts.common')
@section('content')
<div class="row">
  <div class="panel mb25">
      <div class="panel-heading border"> 修改{{$chapter->title}} </div>
      <div class="panel-body">
          <div class="row no-margin">
              <div class="col-lg-12">
                  <form method="post" class="form-horizontal bordered-group" role="form" enctype="multipart/form-data">
                    <input type="text" name="id" value="{{$chapter->id}}" hidden>
                      <div class="form-group">
                          <label class="col-sm-2 control-label">邮件主题：</label>
                          <div class="col-sm-10">
                              <input name="title" class="form-control" value="{{$chapter->title}}"> </div>
                      </div>
                      <div class="form-group">
                          <label class="col-sm-2 control-label">邮件正文：</label>
                          <div class="col-sm-10">
                              <textarea name="content" class="form-control" rows="5">{{$chapter->content}}</textarea>
                          </div>
                      </div>
                      <button type="submit" class="btn btn-info pull-right">提交</button>
                  </form>
              </div>
          </div>
      </div>
  </div>

</div>
@stop
