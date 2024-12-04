from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class M(_message.Message):
    __slots__ = ("memberId", "email", "level")
    MEMBERID_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    memberId: int
    email: str
    level: str
    def __init__(self, memberId: _Optional[int] = ..., email: _Optional[str] = ..., level: _Optional[str] = ...) -> None: ...

class Member(_message.Message):
    __slots__ = ("email", "level", "password")
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    email: str
    level: str
    password: str
    def __init__(self, email: _Optional[str] = ..., level: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class CreateMemberRequest(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: Member
    def __init__(self, member: _Optional[_Union[Member, _Mapping]] = ...) -> None: ...

class UpdateMemberRequest(_message.Message):
    __slots__ = ("member_id", "level", "password", "old_password")
    MEMBER_ID_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    OLD_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    member_id: int
    level: str
    password: str
    old_password: str
    def __init__(self, member_id: _Optional[int] = ..., level: _Optional[str] = ..., password: _Optional[str] = ..., old_password: _Optional[str] = ...) -> None: ...

class DeleteMemberRequest(_message.Message):
    __slots__ = ("member_id", "old_password")
    MEMBER_ID_FIELD_NUMBER: _ClassVar[int]
    OLD_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    member_id: int
    old_password: str
    def __init__(self, member_id: _Optional[int] = ..., old_password: _Optional[str] = ...) -> None: ...

class CreateMemberResponse(_message.Message):
    __slots__ = ("message",)
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...

class UpdateMemberResponse(_message.Message):
    __slots__ = ("message",)
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...

class DeleteMemberResponse(_message.Message):
    __slots__ = ("message",)
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...

class GetMemberByEmailRequest(_message.Message):
    __slots__ = ("email",)
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    email: str
    def __init__(self, email: _Optional[str] = ...) -> None: ...

class GetMemberByEmailResponse(_message.Message):
    __slots__ = ("member_id",)
    MEMBER_ID_FIELD_NUMBER: _ClassVar[int]
    member_id: int
    def __init__(self, member_id: _Optional[int] = ...) -> None: ...

class GetAllMemberRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetAllMemberResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: _containers.RepeatedCompositeFieldContainer[M]
    def __init__(self, member: _Optional[_Iterable[_Union[M, _Mapping]]] = ...) -> None: ...
