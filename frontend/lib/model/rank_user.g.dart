// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'rank_user.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

RankUser _$RankUserFromJson(Map<String, dynamic> json) {
  return RankUser(
    json['name'] as String,
    json['screen_name'] as String,
    json['profile_image'] as String,
    json['point'] as int,
    json['ranking'] as int,
    json['last_ranking'] as int,
  );
}

Map<String, dynamic> _$RankUserToJson(RankUser instance) => <String, dynamic>{
  'name': instance.name,
  'screen_name': instance.screenName,
  'profile_image': instance.profileImage,
  'point': instance.point,
  'ranking': instance.ranking,
  'last_ranking': instance.lastRanking,
};
