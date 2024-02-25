package com.emr.grpcclientdemo.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.61.1)",
    comments = "Source: timeservice.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class TimeServiceGrpc {

  private TimeServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "com.emr.grpcclientdemo.grpc.TimeService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest,
      com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> getNowMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Now",
      requestType = com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest.class,
      responseType = com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest,
      com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> getNowMethod() {
    io.grpc.MethodDescriptor<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest, com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> getNowMethod;
    if ((getNowMethod = TimeServiceGrpc.getNowMethod) == null) {
      synchronized (TimeServiceGrpc.class) {
        if ((getNowMethod = TimeServiceGrpc.getNowMethod) == null) {
          TimeServiceGrpc.getNowMethod = getNowMethod =
              io.grpc.MethodDescriptor.<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest, com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Now"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate.getDefaultInstance()))
              .setSchemaDescriptor(new TimeServiceMethodDescriptorSupplier("Now"))
              .build();
        }
      }
    }
    return getNowMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest,
      com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> getStreamMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Stream",
      requestType = com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest.class,
      responseType = com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest,
      com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> getStreamMethod() {
    io.grpc.MethodDescriptor<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest, com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> getStreamMethod;
    if ((getStreamMethod = TimeServiceGrpc.getStreamMethod) == null) {
      synchronized (TimeServiceGrpc.class) {
        if ((getStreamMethod = TimeServiceGrpc.getStreamMethod) == null) {
          TimeServiceGrpc.getStreamMethod = getStreamMethod =
              io.grpc.MethodDescriptor.<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest, com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Stream"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate.getDefaultInstance()))
              .setSchemaDescriptor(new TimeServiceMethodDescriptorSupplier("Stream"))
              .build();
        }
      }
    }
    return getStreamMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static TimeServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TimeServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TimeServiceStub>() {
        @java.lang.Override
        public TimeServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TimeServiceStub(channel, callOptions);
        }
      };
    return TimeServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static TimeServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TimeServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TimeServiceBlockingStub>() {
        @java.lang.Override
        public TimeServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TimeServiceBlockingStub(channel, callOptions);
        }
      };
    return TimeServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static TimeServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TimeServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TimeServiceFutureStub>() {
        @java.lang.Override
        public TimeServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TimeServiceFutureStub(channel, callOptions);
        }
      };
    return TimeServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void now(com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest request,
        io.grpc.stub.StreamObserver<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getNowMethod(), responseObserver);
    }

    /**
     */
    default void stream(com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest request,
        io.grpc.stub.StreamObserver<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStreamMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service TimeService.
   */
  public static abstract class TimeServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return TimeServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service TimeService.
   */
  public static final class TimeServiceStub
      extends io.grpc.stub.AbstractAsyncStub<TimeServiceStub> {
    private TimeServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TimeServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TimeServiceStub(channel, callOptions);
    }

    /**
     */
    public void now(com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest request,
        io.grpc.stub.StreamObserver<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getNowMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void stream(com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest request,
        io.grpc.stub.StreamObserver<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getStreamMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service TimeService.
   */
  public static final class TimeServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<TimeServiceBlockingStub> {
    private TimeServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TimeServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TimeServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate now(com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getNowMethod(), getCallOptions(), request);
    }

    /**
     */
    public java.util.Iterator<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> stream(
        com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getStreamMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service TimeService.
   */
  public static final class TimeServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<TimeServiceFutureStub> {
    private TimeServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TimeServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TimeServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate> now(
        com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getNowMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_NOW = 0;
  private static final int METHODID_STREAM = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_NOW:
          serviceImpl.now((com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest) request,
              (io.grpc.stub.StreamObserver<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate>) responseObserver);
          break;
        case METHODID_STREAM:
          serviceImpl.stream((com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest) request,
              (io.grpc.stub.StreamObserver<com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getNowMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.NowRequest,
              com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate>(
                service, METHODID_NOW)))
        .addMethod(
          getStreamMethod(),
          io.grpc.stub.ServerCalls.asyncServerStreamingCall(
            new MethodHandlers<
              com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeStreamRequest,
              com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.TimeUpdate>(
                service, METHODID_STREAM)))
        .build();
  }

  private static abstract class TimeServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    TimeServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.emr.grpcclientdemo.grpc.TimeServiceGrpcManager.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("TimeService");
    }
  }

  private static final class TimeServiceFileDescriptorSupplier
      extends TimeServiceBaseDescriptorSupplier {
    TimeServiceFileDescriptorSupplier() {}
  }

  private static final class TimeServiceMethodDescriptorSupplier
      extends TimeServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    TimeServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (TimeServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new TimeServiceFileDescriptorSupplier())
              .addMethod(getNowMethod())
              .addMethod(getStreamMethod())
              .build();
        }
      }
    }
    return result;
  }
}
