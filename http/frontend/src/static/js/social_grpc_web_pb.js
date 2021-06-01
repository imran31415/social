/**
 * @fileoverview gRPC-Web generated client stub for protos
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.protos = require('./social_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.protos.SocialClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.protos.SocialPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.protos.GetUserReq,
 *   !proto.protos.User>}
 */
const methodDescriptor_Social_GetUser = new grpc.web.MethodDescriptor(
  '/protos.Social/GetUser',
  grpc.web.MethodType.UNARY,
  proto.protos.GetUserReq,
  proto.protos.User,
  /**
   * @param {!proto.protos.GetUserReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.User.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.protos.GetUserReq,
 *   !proto.protos.User>}
 */
const methodInfo_Social_GetUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.User,
  /**
   * @param {!proto.protos.GetUserReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.User.deserializeBinary
);


/**
 * @param {!proto.protos.GetUserReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.protos.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.protos.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.protos.SocialClient.prototype.getUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/protos.Social/GetUser',
      request,
      metadata || {},
      methodDescriptor_Social_GetUser,
      callback);
};


/**
 * @param {!proto.protos.GetUserReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.protos.User>}
 *     Promise that resolves to the response
 */
proto.protos.SocialPromiseClient.prototype.getUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/protos.Social/GetUser',
      request,
      metadata || {},
      methodDescriptor_Social_GetUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.protos.CreateUserReq,
 *   !proto.protos.User>}
 */
const methodDescriptor_Social_CreateUser = new grpc.web.MethodDescriptor(
  '/protos.Social/CreateUser',
  grpc.web.MethodType.UNARY,
  proto.protos.CreateUserReq,
  proto.protos.User,
  /**
   * @param {!proto.protos.CreateUserReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.User.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.protos.CreateUserReq,
 *   !proto.protos.User>}
 */
const methodInfo_Social_CreateUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.User,
  /**
   * @param {!proto.protos.CreateUserReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.User.deserializeBinary
);


/**
 * @param {!proto.protos.CreateUserReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.protos.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.protos.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.protos.SocialClient.prototype.createUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/protos.Social/CreateUser',
      request,
      metadata || {},
      methodDescriptor_Social_CreateUser,
      callback);
};


/**
 * @param {!proto.protos.CreateUserReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.protos.User>}
 *     Promise that resolves to the response
 */
proto.protos.SocialPromiseClient.prototype.createUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/protos.Social/CreateUser',
      request,
      metadata || {},
      methodDescriptor_Social_CreateUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.protos.GetPostsReq,
 *   !proto.protos.Posts>}
 */
const methodDescriptor_Social_GetPosts = new grpc.web.MethodDescriptor(
  '/protos.Social/GetPosts',
  grpc.web.MethodType.UNARY,
  proto.protos.GetPostsReq,
  proto.protos.Posts,
  /**
   * @param {!proto.protos.GetPostsReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Posts.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.protos.GetPostsReq,
 *   !proto.protos.Posts>}
 */
const methodInfo_Social_GetPosts = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.Posts,
  /**
   * @param {!proto.protos.GetPostsReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Posts.deserializeBinary
);


/**
 * @param {!proto.protos.GetPostsReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.protos.Posts)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.protos.Posts>|undefined}
 *     The XHR Node Readable Stream
 */
proto.protos.SocialClient.prototype.getPosts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/protos.Social/GetPosts',
      request,
      metadata || {},
      methodDescriptor_Social_GetPosts,
      callback);
};


/**
 * @param {!proto.protos.GetPostsReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.protos.Posts>}
 *     Promise that resolves to the response
 */
proto.protos.SocialPromiseClient.prototype.getPosts =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/protos.Social/GetPosts',
      request,
      metadata || {},
      methodDescriptor_Social_GetPosts);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.protos.CreatePostReq,
 *   !proto.protos.Post>}
 */
const methodDescriptor_Social_CreatePost = new grpc.web.MethodDescriptor(
  '/protos.Social/CreatePost',
  grpc.web.MethodType.UNARY,
  proto.protos.CreatePostReq,
  proto.protos.Post,
  /**
   * @param {!proto.protos.CreatePostReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Post.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.protos.CreatePostReq,
 *   !proto.protos.Post>}
 */
const methodInfo_Social_CreatePost = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.Post,
  /**
   * @param {!proto.protos.CreatePostReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Post.deserializeBinary
);


/**
 * @param {!proto.protos.CreatePostReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.protos.Post)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.protos.Post>|undefined}
 *     The XHR Node Readable Stream
 */
proto.protos.SocialClient.prototype.createPost =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/protos.Social/CreatePost',
      request,
      metadata || {},
      methodDescriptor_Social_CreatePost,
      callback);
};


/**
 * @param {!proto.protos.CreatePostReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.protos.Post>}
 *     Promise that resolves to the response
 */
proto.protos.SocialPromiseClient.prototype.createPost =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/protos.Social/CreatePost',
      request,
      metadata || {},
      methodDescriptor_Social_CreatePost);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.protos.CreateCommentReq,
 *   !proto.protos.Comment>}
 */
const methodDescriptor_Social_CreateComment = new grpc.web.MethodDescriptor(
  '/protos.Social/CreateComment',
  grpc.web.MethodType.UNARY,
  proto.protos.CreateCommentReq,
  proto.protos.Comment,
  /**
   * @param {!proto.protos.CreateCommentReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Comment.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.protos.CreateCommentReq,
 *   !proto.protos.Comment>}
 */
const methodInfo_Social_CreateComment = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.Comment,
  /**
   * @param {!proto.protos.CreateCommentReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Comment.deserializeBinary
);


/**
 * @param {!proto.protos.CreateCommentReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.protos.Comment)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.protos.Comment>|undefined}
 *     The XHR Node Readable Stream
 */
proto.protos.SocialClient.prototype.createComment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/protos.Social/CreateComment',
      request,
      metadata || {},
      methodDescriptor_Social_CreateComment,
      callback);
};


/**
 * @param {!proto.protos.CreateCommentReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.protos.Comment>}
 *     Promise that resolves to the response
 */
proto.protos.SocialPromiseClient.prototype.createComment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/protos.Social/CreateComment',
      request,
      metadata || {},
      methodDescriptor_Social_CreateComment);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.protos.GetCommentsReq,
 *   !proto.protos.Comments>}
 */
const methodDescriptor_Social_GetComments = new grpc.web.MethodDescriptor(
  '/protos.Social/GetComments',
  grpc.web.MethodType.UNARY,
  proto.protos.GetCommentsReq,
  proto.protos.Comments,
  /**
   * @param {!proto.protos.GetCommentsReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Comments.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.protos.GetCommentsReq,
 *   !proto.protos.Comments>}
 */
const methodInfo_Social_GetComments = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.Comments,
  /**
   * @param {!proto.protos.GetCommentsReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Comments.deserializeBinary
);


/**
 * @param {!proto.protos.GetCommentsReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.protos.Comments)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.protos.Comments>|undefined}
 *     The XHR Node Readable Stream
 */
proto.protos.SocialClient.prototype.getComments =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/protos.Social/GetComments',
      request,
      metadata || {},
      methodDescriptor_Social_GetComments,
      callback);
};


/**
 * @param {!proto.protos.GetCommentsReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.protos.Comments>}
 *     Promise that resolves to the response
 */
proto.protos.SocialPromiseClient.prototype.getComments =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/protos.Social/GetComments',
      request,
      metadata || {},
      methodDescriptor_Social_GetComments);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.protos.GetFeedReq,
 *   !proto.protos.Feed>}
 */
const methodDescriptor_Social_GetFeed = new grpc.web.MethodDescriptor(
  '/protos.Social/GetFeed',
  grpc.web.MethodType.UNARY,
  proto.protos.GetFeedReq,
  proto.protos.Feed,
  /**
   * @param {!proto.protos.GetFeedReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Feed.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.protos.GetFeedReq,
 *   !proto.protos.Feed>}
 */
const methodInfo_Social_GetFeed = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.Feed,
  /**
   * @param {!proto.protos.GetFeedReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.Feed.deserializeBinary
);


/**
 * @param {!proto.protos.GetFeedReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.protos.Feed)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.protos.Feed>|undefined}
 *     The XHR Node Readable Stream
 */
proto.protos.SocialClient.prototype.getFeed =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/protos.Social/GetFeed',
      request,
      metadata || {},
      methodDescriptor_Social_GetFeed,
      callback);
};


/**
 * @param {!proto.protos.GetFeedReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.protos.Feed>}
 *     Promise that resolves to the response
 */
proto.protos.SocialPromiseClient.prototype.getFeed =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/protos.Social/GetFeed',
      request,
      metadata || {},
      methodDescriptor_Social_GetFeed);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.protos.CreateFeedItemReq,
 *   !proto.protos.FeedItem>}
 */
const methodDescriptor_Social_CreateFeedItem = new grpc.web.MethodDescriptor(
  '/protos.Social/CreateFeedItem',
  grpc.web.MethodType.UNARY,
  proto.protos.CreateFeedItemReq,
  proto.protos.FeedItem,
  /**
   * @param {!proto.protos.CreateFeedItemReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.FeedItem.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.protos.CreateFeedItemReq,
 *   !proto.protos.FeedItem>}
 */
const methodInfo_Social_CreateFeedItem = new grpc.web.AbstractClientBase.MethodInfo(
  proto.protos.FeedItem,
  /**
   * @param {!proto.protos.CreateFeedItemReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.protos.FeedItem.deserializeBinary
);


/**
 * @param {!proto.protos.CreateFeedItemReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.protos.FeedItem)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.protos.FeedItem>|undefined}
 *     The XHR Node Readable Stream
 */
proto.protos.SocialClient.prototype.createFeedItem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/protos.Social/CreateFeedItem',
      request,
      metadata || {},
      methodDescriptor_Social_CreateFeedItem,
      callback);
};


/**
 * @param {!proto.protos.CreateFeedItemReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.protos.FeedItem>}
 *     Promise that resolves to the response
 */
proto.protos.SocialPromiseClient.prototype.createFeedItem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/protos.Social/CreateFeedItem',
      request,
      metadata || {},
      methodDescriptor_Social_CreateFeedItem);
};


module.exports = proto.protos;

