-include_lib("stdlib/include/assert.hrl").
-define(assertContinueIfMatch(Guard, Expr, Param, With),
        begin
          ((fun () ->
                case (Expr) of
                  Guard -> With(Param);
                  Value -> erlang:error({assertContinueIfMatch,
                                         [{module, ?MODULE},
                                          {line, ?LINE},
                                          {expression, (??Expr)},
                                          {pattern, (??Guard)},
                                          {value, Value}]})
                end
            end)())
        end).

-define(assertCall(Module, Function, Arity, N),
        begin
          ((fun() ->
                __N = (N),
                case meck:num_calls(Module, Function, Arity) of
                  __N -> ok;
                  __V -> erlang:error({assertCall,
                                       [{module, ?MODULE},
                                        {line, ?LINE},
                                        {expression, (??Module) ++ ":" ++ (??Function) ++ "/" ++ (??Arity)},
                                        {expected, __N},
                                        {value, __V}]})
                end
            end)())
        end).

-define(assertCall(Module, Function, N), ?assertCall(Module, Function, '_', N)).
