package com.emr.grpcclientdemo.grpc_service;

import com.google.protobuf.Descriptors;
import lombok.extern.slf4j.Slf4j;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.stereotype.Service;
import com.emr.grpcclientdemo.grpc.TimeServiceGrpc;
import com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager;

@Service
@Slf4j
public class TimeServiceGrpcCaller {

    @GrpcClient("grpc-service")
    private TimeServiceGrpc.TimeServiceBlockingStub stub;

    public void run() {
        var param = TimeServiceGrpcManager.NowRequest.newBuilder()
                .set
                .setLength(2)
                .build();
        TimeServiceGrpcManager.TimeUpdate timeUpdate = stub.now(param);
        log.info(timeUpdate.getTime().getValue());
    }
}
