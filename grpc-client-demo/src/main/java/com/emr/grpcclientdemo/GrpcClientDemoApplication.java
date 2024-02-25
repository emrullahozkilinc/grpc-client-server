package com.emr.grpcclientdemo;

import com.emr.grpcclientdemo.grpc_service.TimeServiceGrpcCaller;
import lombok.RequiredArgsConstructor;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@RequiredArgsConstructor
public class GrpcClientDemoApplication implements CommandLineRunner {

    private final TimeServiceGrpcCaller caller;

    public static void main(String[] args) {
        SpringApplication.run(GrpcClientDemoApplication.class, args);
    }

    @Override
    public void run(String... args) throws Exception {
        caller.run();
    }
}
