// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'ranking.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Ranking _$RankingFromJson(Map<String, dynamic> json) {
  return Ranking(
    (json['rank_users'] as List)
        ?.map((e) =>
    e == null ? null : RankUser.fromJson(e as Map<String, dynamic>))
        ?.toList(),
  );
}

Map<String, dynamic> _$RankingToJson(Ranking instance) => <String, dynamic>{
  'rank_users': instance.rankUsers?.map((e) => e?.toJson())?.toList(),
};
