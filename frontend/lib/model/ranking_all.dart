import 'package:influ_dojo/model/ranking.dart';
import 'package:json_annotation/json_annotation.dart';

part 'ranking_all.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake, explicitToJson: true)
class RankingAll {
  RankingAll(this.dailyWorkRanking, this.dailyResultRanking, this.weeklyWorkRanking,
      this.weeklyResultRanking, this.monthlyWorkRanking, this.monthlyResultRanking);

  Ranking dailyWorkRanking;
  Ranking dailyResultRanking;
  Ranking weeklyWorkRanking;
  Ranking weeklyResultRanking;
  Ranking monthlyWorkRanking;
  Ranking monthlyResultRanking;

  factory RankingAll.fromJson(Map<String, dynamic> json) => _$RankingAllFromJson(json);
  Map<String, dynamic> toJson() => _$RankingAllToJson(this);
}
