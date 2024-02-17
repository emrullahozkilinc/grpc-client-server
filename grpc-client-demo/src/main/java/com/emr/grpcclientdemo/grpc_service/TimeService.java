package com.emr.grpcclientdemo.grpc_service;

import com.google.protobuf.Descriptors;

import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletableFuture;

public interface TimeService {

    Map<Descriptors.FieldDescriptor, Object> Now(NowReq);
    CompletableFuture<List<Map<Descriptors.FieldDescriptor, Object>>> Stream(int authorId) throws InterruptedException;

}
