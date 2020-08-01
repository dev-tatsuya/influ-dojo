// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'ranking_all.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

RankingAll _$RankingAllFromJson(Map<String, dynamic> json) {
  return RankingAll(
    json['daily_work_ranking'] == null
        ? null
        : Ranking.fromJson(json['daily_work_ranking'] as Map<String, dynamic>),
    json['daily_result_ranking'] == null
        ? null
        : Ranking.fromJson(
        json['daily_result_ranking'] as Map<String, dynamic>),
    json['weekly_work_ranking'] == null
        ? null
        : Ranking.fromJson(json['weekly_work_ranking'] as Map<String, dynamic>),
    json['weekly_result_ranking'] == null
        ? null
        : Ranking.fromJson(
        json['weekly_result_ranking'] as Map<String, dynamic>),
    json['monthly_work_ranking'] == null
        ? null
        : Ranking.fromJson(
        json['monthly_work_ranking'] as Map<String, dynamic>),
    json['monthly_result_ranking'] == null
        ? null
        : Ranking.fromJson(
        json['monthly_result_ranking'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$RankingAllToJson(RankingAll instance) =>
    <String, dynamic>{
      'daily_work_ranking': instance.dailyWorkRanking?.toJson(),
      'daily_result_ranking': instance.dailyResultRanking?.toJson(),
      'weekly_work_ranking': instance.weeklyWorkRanking?.toJson(),
      'weekly_result_ranking': instance.weeklyResultRanking?.toJson(),
      'monthly_work_ranking': instance.monthlyWorkRanking?.toJson(),
      'monthly_result_ranking': instance.monthlyResultRanking?.toJson(),
    };
